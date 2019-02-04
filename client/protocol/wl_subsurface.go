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

// WlSubsurfaceError is for .
//
//
type WlSubsurfaceError int

const (
	WlSubsurfaceErrorBadSurface WlSubsurfaceError = 0 // wl_surface is not a sibling or the parent
)

/*
 * EVENT HANDLER TYPES
 */

/*
 * TYPE
 */
// WlSubsurface is sub-surface interface to a wl_surface.
//
// An additional interface to a wl_surface object, which has been
// made a sub-surface. A sub-surface has one parent surface. A
// sub-surface's size and position are not limited to that of the parent.
// Particularly, a sub-surface is not automatically clipped to its
// parent's area.
//
// A sub-surface becomes mapped, when a non-NULL wl_buffer is applied
// and the parent surface is mapped. The order of which one happens
// first is irrelevant. A sub-surface is hidden if the parent becomes
// hidden, or if a NULL wl_buffer is applied. These rules apply
// recursively through the tree of surfaces.
//
// The behaviour of a wl_surface.commit request on a sub-surface
// depends on the sub-surface's mode. The possible modes are
// synchronized and desynchronized, see methods
// wl_subsurface.set_sync and wl_subsurface.set_desync. Synchronized
// mode caches the wl_surface state to be applied when the parent's
// state gets applied, and desynchronized mode applies the pending
// wl_surface state directly. A sub-surface is initially in the
// synchronized mode.
//
// Sub-surfaces have also other kind of state, which is managed by
// wl_subsurface requests, as opposed to wl_surface requests. This
// state includes the sub-surface position relative to the parent
// surface (wl_subsurface.set_position), and the stacking order of
// the parent and its sub-surfaces (wl_subsurface.place_above and
// .place_below). This state is applied when the parent surface's
// wl_surface state is applied, regardless of the sub-surface's mode.
// As the exception, set_sync and set_desync are effective immediately.
//
// The main surface can be thought to be always in desynchronized mode,
// since it does not have a parent in the sub-surfaces sense.
//
// Even if a sub-surface is in desynchronized mode, it will behave as
// in synchronized mode, if its parent surface behaves as in
// synchronized mode. This rule is applied recursively throughout the
// tree of surfaces. This means, that one can set a sub-surface into
// synchronized mode, and then assume that all its child and grand-child
// sub-surfaces are synchronized, too, without explicitly setting them.
//
// If the wl_surface associated with the wl_subsurface is destroyed, the
// wl_subsurface object becomes inert. Note, that destroying either object
// takes effect immediately. If you need to synchronize the removal
// of a sub-surface to the parent surface update, unmap the sub-surface
// first by attaching a NULL wl_buffer, update parent, and then destroy
// the sub-surface.
//
// If the parent wl_surface object is destroyed, the sub-surface is
// unmapped.
type WlSubsurface struct {
	Base
}

// NewWlSubsurface creates a WlSubsurface object.
func NewWlSubsurface(c *wire.Conn) *WlSubsurface {
	return NewWlSubsurfaceWithID(c, c.NewID())
}

// NewWlSubsurfaceWithID creates a WlSubsurface object with a given id.
func NewWlSubsurfaceWithID(c *wire.Conn, id wire.ID) *WlSubsurface {
	o := &WlSubsurface{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Destroy is for remove sub-surface interface
//
// The sub-surface interface is removed from the wl_surface object
// that was turned into a sub-surface with a
// wl_subcompositor.get_subsurface request. The wl_surface's association
// to the parent is deleted, and the wl_surface loses its role as
// a sub-surface. The wl_surface is unmapped immediately.
func (o *WlSubsurface) Destroy() error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetPosition is for reposition the sub-surface
//
// This schedules a sub-surface position change.
// The sub-surface will be moved so that its origin (top left
// corner pixel) will be at the location x, y of the parent surface
// coordinate system. The coordinates are not restricted to the parent
// surface area. Negative values are allowed.
//
// The scheduled coordinates will take effect whenever the state of the
// parent surface is applied. When this happens depends on whether the
// parent surface is in synchronized mode or not. See
// wl_subsurface.set_sync and wl_subsurface.set_desync for details.
//
// If more than one set_position request is invoked by the client before
// the commit of the parent surface, the position of a new request always
// replaces the scheduled position from any previous request.
//
// The initial position is 0, 0.
func (o *WlSubsurface) SetPosition(x wire.Int, y wire.Int) error {
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

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// PlaceAbove is for restack the sub-surface
//
// This sub-surface is taken from the stack, and put back just
// above the reference surface, changing the z-order of the sub-surfaces.
// The reference surface must be one of the sibling surfaces, or the
// parent surface. Using any other surface, including this sub-surface,
// will cause a protocol error.
//
// The z-order is double-buffered. Requests are handled in order and
// applied immediately to a pending state. The final pending state is
// copied to the active state the next time the state of the parent
// surface is applied. When this happens depends on whether the parent
// surface is in synchronized mode or not. See wl_subsurface.set_sync and
// wl_subsurface.set_desync for details.
//
// A new sub-surface is initially added as the top-most in the stack
// of its siblings and parent.
func (o *WlSubsurface) PlaceAbove(sibling *WlSurface) error {
	msg, err := wire.NewMessage(o.ID(), 2)
	if err != nil {
		return err
	}

	if err = msg.Write(sibling.ID()); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// PlaceBelow is for restack the sub-surface
//
// The sub-surface is placed just below the reference surface.
// See wl_subsurface.place_above.
func (o *WlSubsurface) PlaceBelow(sibling *WlSurface) error {
	msg, err := wire.NewMessage(o.ID(), 3)
	if err != nil {
		return err
	}

	if err = msg.Write(sibling.ID()); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetSync is for set sub-surface to synchronized mode
//
// Change the commit behaviour of the sub-surface to synchronized
// mode, also described as the parent dependent mode.
//
// In synchronized mode, wl_surface.commit on a sub-surface will
// accumulate the committed state in a cache, but the state will
// not be applied and hence will not change the compositor output.
// The cached state is applied to the sub-surface immediately after
// the parent surface's state is applied. This ensures atomic
// updates of the parent and all its synchronized sub-surfaces.
// Applying the cached state will invalidate the cache, so further
// parent surface commits do not (re-)apply old state.
//
// See wl_subsurface for the recursive effect of this mode.
func (o *WlSubsurface) SetSync() error {
	msg, err := wire.NewMessage(o.ID(), 4)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetDesync is for set sub-surface to desynchronized mode
//
// Change the commit behaviour of the sub-surface to desynchronized
// mode, also described as independent or freely running mode.
//
// In desynchronized mode, wl_surface.commit on a sub-surface will
// apply the pending state directly, without caching, as happens
// normally with a wl_surface. Calling wl_surface.commit on the
// parent surface has no effect on the sub-surface's wl_surface
// state. This mode allows a sub-surface to be updated on its own.
//
// If cached state exists when wl_surface.commit is called in
// desynchronized mode, the pending state is added to the cached
// state, and applied as a whole. This invalidates the cache.
//
// Note: even if a sub-surface is set to desynchronized, a parent
// sub-surface may override it to behave as synchronized. For details,
// see wl_subsurface.
//
// If a surface's parent surface behaves as desynchronized, then
// the cached state is applied on set_desync.
func (o *WlSubsurface) SetDesync() error {
	msg, err := wire.NewMessage(o.ID(), 5)
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

// ServeMessage is a multiplexer for a message.
func (o *WlSubsurface) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	default:
		return fmt.Errorf("WlSubsurface: unhandled message(%v)", msg)
	}
}
