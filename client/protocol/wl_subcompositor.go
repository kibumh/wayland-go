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

// WlSubcompositorError is for .
//
//
type WlSubcompositorError int

const (
	WlSubcompositorErrorBadSurface WlSubcompositorError = 0 // the to-be sub-surface is invalid
)

/*
 * EVENT HANDLER TYPES
 */

/*
 * TYPE
 */
// WlSubcompositor is sub-surface compositing.
//
// The global interface exposing sub-surface compositing capabilities.
// A wl_surface, that has sub-surfaces associated, is called the
// parent surface. Sub-surfaces can be arbitrarily nested and create
// a tree of sub-surfaces.
//
// The root surface in a tree of sub-surfaces is the main
// surface. The main surface cannot be a sub-surface, because
// sub-surfaces must always have a parent.
//
// A main surface with its sub-surfaces forms a (compound) window.
// For window management purposes, this set of wl_surface objects is
// to be considered as a single window, and it should also behave as
// such.
//
// The aim of sub-surfaces is to offload some of the compositing work
// within a window from clients to the compositor. A prime example is
// a video player with decorations and video in separate wl_surface
// objects. This should allow the compositor to pass YUV video buffer
// processing to dedicated overlay hardware when possible.
type WlSubcompositor struct {
	Base
}

// NewWlSubcompositor creates a WlSubcompositor object.
func NewWlSubcompositor(c *wire.Conn) *WlSubcompositor {
	return NewWlSubcompositorWithID(c, c.NewID())
}

// NewWlSubcompositorWithID creates a WlSubcompositor object with a given id.
func NewWlSubcompositorWithID(c *wire.Conn, id wire.ID) *WlSubcompositor {
	o := &WlSubcompositor{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Destroy is for unbind from the subcompositor interface
//
// Informs the server that the client will not be using this
// protocol object anymore. This does not affect any other
// objects, wl_subsurface objects included.
func (o *WlSubcompositor) Destroy() error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// GetSubsurface is for give a surface the role sub-surface
//
// Create a sub-surface interface for the given surface, and
// associate it with the given parent surface. This turns a
// plain wl_surface into a sub-surface.
//
// The to-be sub-surface must not already have another role, and it
// must not have an existing wl_subsurface object. Otherwise a protocol
// error is raised.
//
// Adding sub-surfaces to a parent is a double-buffered operation on the
// parent (see wl_surface.commit). The effect of adding a sub-surface
// becomes visible on the next time the state of the parent surface is
// applied.
//
// This request modifies the behaviour of wl_surface.commit request on
// the sub-surface, see the documentation on wl_subsurface interface.
func (o *WlSubcompositor) GetSubsurface(id *WlSubsurface, surface *WlSurface, parent *WlSurface) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(id.ID()); err != nil {
		return err
	}

	if err = msg.Write(surface.ID()); err != nil {
		return err
	}

	if err = msg.Write(parent.ID()); err != nil {
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
func (o *WlSubcompositor) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	default:
		return fmt.Errorf("WlSubcompositor: unhandled message(%v)", msg)
	}
}
