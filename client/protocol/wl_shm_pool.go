/*
 * THIS CODE IS GENERATED. DO NOT EDIT BY HANDS.
 */

package protocol

import (
	"fmt"

	"github.com/kibumh/wayland-go/wire"
)

/*
 * ENUMS
 */

/*
 * EVENT HANDLER TYPES
 */

/*
 * TYPE
 */
// WlShmPool is a shared memory pool.
//
// The wl_shm_pool object encapsulates a piece of memory shared
// between the compositor and client.  Through the wl_shm_pool
// object, the client can allocate shared memory wl_buffer objects.
// All objects created through the same pool share the same
// underlying mapped memory. Reusing the mapped memory avoids the
// setup/teardown overhead and is useful when interactively resizing
// a surface or for many small buffers.
type WlShmPool struct {
	Base
}

// NewWlShmPool creates a WlShmPool object.
func NewWlShmPool(c *wire.Conn) *WlShmPool {
	return NewWlShmPoolWithID(c, c.NewID())
}

// NewWlShmPoolWithID creates a WlShmPool object with a given id.
func NewWlShmPoolWithID(c *wire.Conn, id wire.ID) *WlShmPool {
	o := &WlShmPool{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// CreateBuffer is for create a buffer from the pool
//
// Create a wl_buffer object from the pool.
//
// The buffer is created offset bytes into the pool and has
// width and height as specified.  The stride argument specifies
// the number of bytes from the beginning of one row to the beginning
// of the next.  The format is the pixel format of the buffer and
// must be one of those advertised through the wl_shm.format event.
//
// A buffer will keep a reference to the pool it was created from
// so it is valid to destroy the pool immediately after creating
// a buffer from it.
func (o *WlShmPool) CreateBuffer(id *WlBuffer, offset wire.Int, width wire.Int, height wire.Int, stride wire.Int, format wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(id.ID()); err != nil {
		return err
	}

	if err = msg.Write(offset); err != nil {
		return err
	}

	if err = msg.Write(width); err != nil {
		return err
	}

	if err = msg.Write(height); err != nil {
		return err
	}

	if err = msg.Write(stride); err != nil {
		return err
	}

	if err = msg.Write(format); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Destroy is for destroy the pool
//
// Destroy the shared memory pool.
//
// The mmapped memory will be released when all
// buffers that have been created from this pool
// are gone.
func (o *WlShmPool) Destroy() error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Resize is for change the size of the pool mapping
//
// This request will cause the server to remap the backing memory
// for the pool from the file descriptor passed when the pool was
// created, but using the new size.  This request can only be
// used to make the pool bigger.
func (o *WlShmPool) Resize(size wire.Int) error {
	msg, err := wire.NewMessage(o.ID(), 2)
	if err != nil {
		return err
	}

	if err = msg.Write(size); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

/*
 * EVENTS
 */

// ServeMessage is a multiplexer for a message.
func (o *WlShmPool) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	default:
		return fmt.Errorf("WlShmPool: unhandled message(%v)", msg)
	}
}
