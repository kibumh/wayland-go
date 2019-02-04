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
// WlCompositor is the compositor singleton.
//
// A compositor.  This object is a singleton global.  The
// compositor is in charge of combining the contents of multiple
// surfaces into one displayable output.
type WlCompositor struct {
	Base
}

// NewWlCompositor creates a WlCompositor object.
func NewWlCompositor(c *wire.Conn) *WlCompositor {
	return NewWlCompositorWithID(c, c.NewID())
}

// NewWlCompositorWithID creates a WlCompositor object with a given id.
func NewWlCompositorWithID(c *wire.Conn, id wire.ID) *WlCompositor {
	o := &WlCompositor{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// CreateSurface is for create new surface
//
// Ask the compositor to create a new surface.
func (o *WlCompositor) CreateSurface(id *WlSurface) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(id.ID()); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// CreateRegion is for create new region
//
// Ask the compositor to create a new region.
func (o *WlCompositor) CreateRegion(id *WlRegion) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(id.ID()); err != nil {
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
func (o *WlCompositor) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	default:
		return fmt.Errorf("WlCompositor: unhandled message(%v)", msg)
	}
}
