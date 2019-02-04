package wire

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Message struct {
	id      ID
	Opcode  int
	Payload *bytes.Buffer
}

var hostEndian = binary.LittleEndian

func NewMessage(id ID, opcode int) (*Message, error) {
	return &Message{id: id, Opcode: opcode, Payload: &bytes.Buffer{}}, nil
}

func (m *Message) Write(args ...interface{}) error {
	for _, arg := range args {
		switch v := arg.(type) {
		case ID:
			if err := binary.Write(m.Payload, hostEndian, v); err != nil {
				return errors.Wrap(err, "msg.Write: write failed")
			}
		default:
			return fmt.Errorf("msg.Write: not implemented type(%v)", v)
		}
	}
	return nil
}

type MessageReader struct {
	b *bytes.Buffer
}

func NewReader(b *bytes.Buffer) *MessageReader {
	return &MessageReader{b: b}
}

func (r *MessageReader) ReadUint() (Uint, error) {
	var u Uint
	if err := binary.Read(r.b, hostEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func (r *MessageReader) ReadInt() (Int, error) {
	var u Int
	if err := binary.Read(r.b, hostEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func (r *MessageReader) ReadID() (ID, error) {
	var u ID
	if err := binary.Read(r.b, hostEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func (r *MessageReader) ReadString() (String, error) {
	var len int32
	if err := binary.Read(r.b, hostEndian, &len); err != nil {
		return "", err
	}
	adjlen := (len + 3) / 4 * 4
	buf := make([]byte, adjlen)
	if _, err := io.ReadFull(r.b, buf); err != nil {
		return "", err
	}
	return String(buf[:len-1]), nil
}

func (r *MessageReader) ReadFixed() (Fixed, error) {
	return 0, nil
}

func (r *MessageReader) ReadFD() (FD, error) {
	return 0, nil
}

func (r *MessageReader) ReadArray() ([]byte, error) {
	return nil, nil
}
