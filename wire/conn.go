package wire

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"sync"

	"github.com/pkg/errors"
)

// Conn represents a client connection.
type Conn struct {
	uconn *net.UnixConn

	mu     sync.Mutex
	maxID  ID
	objMap map[ID]Object
}

// Dial dials to a wayland remote server at the given address.
func Dial(addr string) (*Conn, error) {
	uaddr, err := net.ResolveUnixAddr("unix", addr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve '%s'", addr)
	}

	uconn, err := net.DialUnix("unix", nil, uaddr)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to dial to %s", addr)
	}

	c := &Conn{uconn: uconn,
		objMap: make(map[ID]Object),
	}

	// go func() {
	// loop:
	// 	for i := 0; ; i++ {
	// 		select {
	// 		case <-c.done:
	// 			break loop
	// 		case c.idchan <- ID(i):
	// 		}
	// 	}
	// 	close(c.idchan)
	// }()
	go c.recvLoop()

	return c, nil
}

// Close closes the connection.
func (c *Conn) Close() error {
	// c.done <- struct{}{}
	return c.uconn.Close()
}

// NewID return a new id.
func (c *Conn) NewID() ID {
	c.mu.Lock()
	c.maxID++
	id := c.maxID
	c.mu.Unlock()
	return id
}

func (c *Conn) RegisterObject(o Object) {
	c.mu.Lock()
	c.objMap[o.ID()] = o
	c.mu.Unlock()
}

func (c *Conn) GetObject(id ID) (Object, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	o, ok := c.objMap[id]
	return o, ok
}

func (c *Conn) Write(msg *Message) error {
	var buf bytes.Buffer
	if err := binary.Write(&buf, hostEndian, msg.id); err != nil {
		return err
	}
	var data uint32
	data = uint32(msg.Payload.Len() + 8)
	data <<= 16
	data |= uint32(msg.Opcode)
	if err := binary.Write(&buf, hostEndian, data); err != nil {
		return err
	}
	if _, err := buf.ReadFrom(msg.Payload); err != nil {
		return nil
	}
	log.Printf("conn.write: %v", buf.Bytes)

	_, _, err := c.uconn.WriteMsgUnix(buf.Bytes(), nil, nil)
	return err
}

func (c *Conn) Read(msg *Message) error {
	var id ID
	if err := binary.Read(c.uconn, hostEndian, &id); err != nil {
		return err
	}
	var data uint32
	if err := binary.Read(c.uconn, hostEndian, &data); err != nil {
		return err
	}
	payload := make([]byte, int((data>>16)&0xffff)-8)
	if err := binary.Read(c.uconn, hostEndian, &payload); err != nil {
		return err
	}

	msg.id = id
	msg.Opcode = int(data & 0xffff)
	msg.Payload = bytes.NewBuffer(payload)
	log.Printf("conn.read: %+v, %d", msg, len(payload))
	return nil
}

func (c *Conn) recvLoop() error {
	for {
		var msg Message
		if err := c.Read(&msg); err != nil {
			return nil
		}
		o, ok := c.GetObject(msg.id)
		if !ok {
			log.Printf("mux.recvLoop: no object found for id(%v)", msg.id)
			continue
		}
		log.Printf("msg to object(%+v)", o)
		o.ServeMessage(&msg)
	}
}
