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

// WlBufferReleaseHandler is a handler for compositor releases buffer.
//
// Sent when this wl_buffer is no longer used by the compositor.
// The client is now free to reuse or destroy this buffer and its
// backing storage.
//
// If a client receives a release event before the frame callback
// requested in the same wl_surface.commit that attaches this
// wl_buffer to a surface, then the client is immediately free to
// reuse the buffer and its backing storage, and does not need a
// second buffer for the next surface content update. Typically
// this is possible, when the compositor maintains a copy of the
// wl_surface contents, e.g. as a GL texture. This is an important
// optimization for GL(ES) compositors with wl_shm clients.
type WlBufferReleaseHandler func() error

/*
 * TYPE
 */
// WlBuffer is content for a wl_surface.
//
// A buffer provides the content for a wl_surface. Buffers are
// created through factory interfaces such as wl_drm, wl_shm or
// similar. It has a width and a height and can be attached to a
// wl_surface, but the mechanism by which a client provides and
// updates the contents is defined by the buffer factory interface.
type WlBuffer struct {
	Base

	releaseHandler WlBufferReleaseHandler
}

// NewWlBuffer creates a WlBuffer object.
func NewWlBuffer(c *wire.Conn) *WlBuffer {
	return NewWlBufferWithID(c, c.NewID())
}

// NewWlBufferWithID creates a WlBuffer object with a given id.
func NewWlBufferWithID(c *wire.Conn, id wire.ID) *WlBuffer {
	o := &WlBuffer{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Destroy is for destroy a buffer
//
// Destroy a buffer. If and how you need to release the backing
// storage is defined by the buffer factory interface.
//
// For possible side-effects to a surface, see wl_surface.attach.
func (o *WlBuffer) Destroy() error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
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

// ServeRelease is for compositor releases buffer.
//
// Sent when this wl_buffer is no longer used by the compositor.
// The client is now free to reuse or destroy this buffer and its
// backing storage.
//
// If a client receives a release event before the frame callback
// requested in the same wl_surface.commit that attaches this
// wl_buffer to a surface, then the client is immediately free to
// reuse the buffer and its backing storage, and does not need a
// second buffer for the next surface content update. Typically
// this is possible, when the compositor maintains a copy of the
// wl_surface contents, e.g. as a GL texture. This is an important
// optimization for GL(ES) compositors with wl_shm clients.
func (o *WlBuffer) ServeRelease() error {
	if o.releaseHandler == nil {
		return nil
	}
	return o.releaseHandler()
}

// HandleRelease registers a handler for a Release event.
func (o *WlBuffer) HandleRelease(h WlBufferReleaseHandler) {
	o.releaseHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlBuffer) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.releaseHandler == nil {
			return nil
		}

		return o.releaseHandler()

	default:
		return fmt.Errorf("WlBuffer: unhandled message(%v)", msg)
	}
}
