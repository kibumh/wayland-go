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
// WlRegion is region interface.
//
// A region object describes an area.
//
// Region objects are used to describe the opaque and input
// regions of a surface.
type WlRegion struct {
	Base
}

// NewWlRegion creates a WlRegion object.
func NewWlRegion(c *wire.Conn) *WlRegion {
	return NewWlRegionWithID(c, c.NewID())
}

// NewWlRegionWithID creates a WlRegion object with a given id.
func NewWlRegionWithID(c *wire.Conn, id wire.ID) *WlRegion {
	o := &WlRegion{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Destroy is for destroy region
//
// Destroy the region.  This will invalidate the object ID.
func (o *WlRegion) Destroy() error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Add is for add rectangle to region
//
// Add the specified rectangle to the region.
func (o *WlRegion) Add(x wire.Int, y wire.Int, width wire.Int, height wire.Int) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(x); err != nil {
		return err
	}

	if err = msg.Write(y); err != nil {
		return err
	}

	if err = msg.Write(width); err != nil {
		return err
	}

	if err = msg.Write(height); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Subtract is for subtract rectangle from region
//
// Subtract the specified rectangle from the region.
func (o *WlRegion) Subtract(x wire.Int, y wire.Int, width wire.Int, height wire.Int) error {
	msg, err := wire.NewMessage(o.ID(), 2)
	if err != nil {
		return err
	}

	if err = msg.Write(x); err != nil {
		return err
	}

	if err = msg.Write(y); err != nil {
		return err
	}

	if err = msg.Write(width); err != nil {
		return err
	}

	if err = msg.Write(height); err != nil {
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
func (o *WlRegion) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	default:
		return fmt.Errorf("WlRegion: unhandled message(%v)", msg)
	}
}
