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

// WlShellError is for .
//
//
type WlShellError int

const (
	WlShellErrorRole WlShellError = 0 // given wl_surface has another role
)

/*
 * EVENT HANDLER TYPES
 */

/*
 * TYPE
 */
// WlShell is create desktop-style surfaces.
//
// This interface is implemented by servers that provide
// desktop-style user interfaces.
//
// It allows clients to associate a wl_shell_surface with
// a basic surface.
//
// Note! This protocol is deprecated and not intended for production use.
// For desktop-style user interfaces, use xdg_shell.
type WlShell struct {
	Base
}

// NewWlShell creates a WlShell object.
func NewWlShell(c *wire.Conn) *WlShell {
	return NewWlShellWithID(c, c.NewID())
}

// NewWlShellWithID creates a WlShell object with a given id.
func NewWlShellWithID(c *wire.Conn, id wire.ID) *WlShell {
	o := &WlShell{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// GetShellSurface is for create a shell surface from a surface
//
// Create a shell surface for an existing surface. This gives
// the wl_surface the role of a shell surface. If the wl_surface
// already has another role, it raises a protocol error.
//
// Only one shell surface can be associated with a given surface.
func (o *WlShell) GetShellSurface(id *WlShellSurface, surface *WlSurface) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(id.ID()); err != nil {
		return err
	}

	if err = msg.Write(surface.ID()); err != nil {
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
func (o *WlShell) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	default:
		return fmt.Errorf("WlShell: unhandled message(%v)", msg)
	}
}
