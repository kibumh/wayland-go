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

// WlShellSurfaceResize is for edge values for resizing.
//
// These values are used to indicate which edge of a surface
// is being dragged in a resize operation. The server may
// use this information to adapt its behavior, e.g. choose
// an appropriate cursor image.
type WlShellSurfaceResize int

const (
	WlShellSurfaceResizeNone        WlShellSurfaceResize = 0  // no edge
	WlShellSurfaceResizeTop         WlShellSurfaceResize = 1  // top edge
	WlShellSurfaceResizeBottom      WlShellSurfaceResize = 2  // bottom edge
	WlShellSurfaceResizeLeft        WlShellSurfaceResize = 4  // left edge
	WlShellSurfaceResizeTopLeft     WlShellSurfaceResize = 5  // top and left edges
	WlShellSurfaceResizeBottomLeft  WlShellSurfaceResize = 6  // bottom and left edges
	WlShellSurfaceResizeRight       WlShellSurfaceResize = 8  // right edge
	WlShellSurfaceResizeTopRight    WlShellSurfaceResize = 9  // top and right edges
	WlShellSurfaceResizeBottomRight WlShellSurfaceResize = 10 // bottom and right edges
)

// WlShellSurfaceTransient is for details of transient behaviour.
//
// These flags specify details of the expected behaviour
// of transient surfaces. Used in the set_transient request.
type WlShellSurfaceTransient int

const (
	WlShellSurfaceTransientInactive WlShellSurfaceTransient = 0x1 // do not set keyboard focus
)

// WlShellSurfaceFullscreenMethod is for different method to set the surface fullscreen.
//
// Hints to indicate to the compositor how to deal with a conflict
// between the dimensions of the surface and the dimensions of the
// output. The compositor is free to ignore this parameter.
type WlShellSurfaceFullscreenMethod int

const (
	WlShellSurfaceFullscreenMethodDefault WlShellSurfaceFullscreenMethod = 0 // no preference, apply default policy
	WlShellSurfaceFullscreenMethodScale   WlShellSurfaceFullscreenMethod = 1 // scale, preserve the surface's aspect ratio and center on output
	WlShellSurfaceFullscreenMethodDriver  WlShellSurfaceFullscreenMethod = 2 // switch output mode to the smallest mode that can fit the surface, add black borders to compensate size mismatch
	WlShellSurfaceFullscreenMethodFill    WlShellSurfaceFullscreenMethod = 3 // no upscaling, center on output and add black borders to compensate size mismatch
)

/*
 * EVENT HANDLER TYPES
 */

// WlShellSurfacePingHandler is a handler for ping client.
//
// Ping a client to check if it is receiving events and sending
// requests. A client is expected to reply with a pong request.
type WlShellSurfacePingHandler func(serial wire.Uint) error

// WlShellSurfaceConfigureHandler is a handler for suggest resize.
//
// The configure event asks the client to resize its surface.
//
// The size is a hint, in the sense that the client is free to
// ignore it if it doesn't resize, pick a smaller size (to
// satisfy aspect ratio or resize in steps of NxM pixels).
//
// The edges parameter provides a hint about how the surface
// was resized. The client may use this information to decide
// how to adjust its content to the new size (e.g. a scrolling
// area might adjust its content position to leave the viewable
// content unmoved).
//
// The client is free to dismiss all but the last configure
// event it received.
//
// The width and height arguments specify the size of the window
// in surface-local coordinates.
type WlShellSurfaceConfigureHandler func(edges wire.Uint, width wire.Int, height wire.Int) error

// WlShellSurfacePopupDoneHandler is a handler for popup interaction is done.
//
// The popup_done event is sent out when a popup grab is broken,
// that is, when the user clicks a surface that doesn't belong
// to the client owning the popup surface.
type WlShellSurfacePopupDoneHandler func() error

/*
 * TYPE
 */
// WlShellSurface is desktop-style metadata interface.
//
// An interface that may be implemented by a wl_surface, for
// implementations that provide a desktop-style user interface.
//
// It provides requests to treat surfaces like toplevel, fullscreen
// or popup windows, move, resize or maximize them, associate
// metadata like title and class, etc.
//
// On the server side the object is automatically destroyed when
// the related wl_surface is destroyed. On the client side,
// wl_shell_surface_destroy() must be called before destroying
// the wl_surface object.
type WlShellSurface struct {
	Base

	pingHandler      WlShellSurfacePingHandler
	configureHandler WlShellSurfaceConfigureHandler
	popupDoneHandler WlShellSurfacePopupDoneHandler
}

// NewWlShellSurface creates a WlShellSurface object.
func NewWlShellSurface(c *wire.Conn) *WlShellSurface {
	return NewWlShellSurfaceWithID(c, c.NewID())
}

// NewWlShellSurfaceWithID creates a WlShellSurface object with a given id.
func NewWlShellSurfaceWithID(c *wire.Conn, id wire.ID) *WlShellSurface {
	o := &WlShellSurface{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Pong is for respond to a ping event
//
// A client must respond to a ping event with a pong request or
// the client may be deemed unresponsive.
func (o *WlShellSurface) Pong(serial wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(serial); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Move is for start an interactive move
//
// Start a pointer-driven move of the surface.
//
// This request must be used in response to a button press event.
// The server may ignore move requests depending on the state of
// the surface (e.g. fullscreen or maximized).
func (o *WlShellSurface) Move(seat *WlSeat, serial wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(seat.ID()); err != nil {
		return err
	}

	if err = msg.Write(serial); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Resize is for start an interactive resize
//
// Start a pointer-driven resizing of the surface.
//
// This request must be used in response to a button press event.
// The server may ignore resize requests depending on the state of
// the surface (e.g. fullscreen or maximized).
func (o *WlShellSurface) Resize(seat *WlSeat, serial wire.Uint, edges wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 2)
	if err != nil {
		return err
	}

	if err = msg.Write(seat.ID()); err != nil {
		return err
	}

	if err = msg.Write(serial); err != nil {
		return err
	}

	if err = msg.Write(edges); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetToplevel is for make the surface a toplevel surface
//
// Map the surface as a toplevel surface.
//
// A toplevel surface is not fullscreen, maximized or transient.
func (o *WlShellSurface) SetToplevel() error {
	msg, err := wire.NewMessage(o.ID(), 3)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetTransient is for make the surface a transient surface
//
// Map the surface relative to an existing surface.
//
// The x and y arguments specify the location of the upper left
// corner of the surface relative to the upper left corner of the
// parent surface, in surface-local coordinates.
//
// The flags argument controls details of the transient behaviour.
func (o *WlShellSurface) SetTransient(parent *WlSurface, x wire.Int, y wire.Int, flags wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 4)
	if err != nil {
		return err
	}

	if err = msg.Write(parent.ID()); err != nil {
		return err
	}

	if err = msg.Write(x); err != nil {
		return err
	}

	if err = msg.Write(y); err != nil {
		return err
	}

	if err = msg.Write(flags); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetFullscreen is for make the surface a fullscreen surface
//
// Map the surface as a fullscreen surface.
//
// If an output parameter is given then the surface will be made
// fullscreen on that output. If the client does not specify the
// output then the compositor will apply its policy - usually
// choosing the output on which the surface has the biggest surface
// area.
//
// The client may specify a method to resolve a size conflict
// between the output size and the surface size - this is provided
// through the method parameter.
//
// The framerate parameter is used only when the method is set
// to "driver", to indicate the preferred framerate. A value of 0
// indicates that the client does not care about framerate.  The
// framerate is specified in mHz, that is framerate of 60000 is 60Hz.
//
// A method of "scale" or "driver" implies a scaling operation of
// the surface, either via a direct scaling operation or a change of
// the output mode. This will override any kind of output scaling, so
// that mapping a surface with a buffer size equal to the mode can
// fill the screen independent of buffer_scale.
//
// A method of "fill" means we don't scale up the buffer, however
// any output scale is applied. This means that you may run into
// an edge case where the application maps a buffer with the same
// size of the output mode but buffer_scale 1 (thus making a
// surface larger than the output). In this case it is allowed to
// downscale the results to fit the screen.
//
// The compositor must reply to this request with a configure event
// with the dimensions for the output on which the surface will
// be made fullscreen.
func (o *WlShellSurface) SetFullscreen(method wire.Uint, framerate wire.Uint, output *WlOutput) error {
	msg, err := wire.NewMessage(o.ID(), 5)
	if err != nil {
		return err
	}

	if err = msg.Write(method); err != nil {
		return err
	}

	if err = msg.Write(framerate); err != nil {
		return err
	}

	if err = msg.Write(output.ID()); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetPopup is for make the surface a popup surface
//
// Map the surface as a popup.
//
// A popup surface is a transient surface with an added pointer
// grab.
//
// An existing implicit grab will be changed to owner-events mode,
// and the popup grab will continue after the implicit grab ends
// (i.e. releasing the mouse button does not cause the popup to
// be unmapped).
//
// The popup grab continues until the window is destroyed or a
// mouse button is pressed in any other client's window. A click
// in any of the client's surfaces is reported as normal, however,
// clicks in other clients' surfaces will be discarded and trigger
// the callback.
//
// The x and y arguments specify the location of the upper left
// corner of the surface relative to the upper left corner of the
// parent surface, in surface-local coordinates.
func (o *WlShellSurface) SetPopup(seat *WlSeat, serial wire.Uint, parent *WlSurface, x wire.Int, y wire.Int, flags wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 6)
	if err != nil {
		return err
	}

	if err = msg.Write(seat.ID()); err != nil {
		return err
	}

	if err = msg.Write(serial); err != nil {
		return err
	}

	if err = msg.Write(parent.ID()); err != nil {
		return err
	}

	if err = msg.Write(x); err != nil {
		return err
	}

	if err = msg.Write(y); err != nil {
		return err
	}

	if err = msg.Write(flags); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetMaximized is for make the surface a maximized surface
//
// Map the surface as a maximized surface.
//
// If an output parameter is given then the surface will be
// maximized on that output. If the client does not specify the
// output then the compositor will apply its policy - usually
// choosing the output on which the surface has the biggest surface
// area.
//
// The compositor will reply with a configure event telling
// the expected new surface size. The operation is completed
// on the next buffer attach to this surface.
//
// A maximized surface typically fills the entire output it is
// bound to, except for desktop elements such as panels. This is
// the main difference between a maximized shell surface and a
// fullscreen shell surface.
//
// The details depend on the compositor implementation.
func (o *WlShellSurface) SetMaximized(output *WlOutput) error {
	msg, err := wire.NewMessage(o.ID(), 7)
	if err != nil {
		return err
	}

	if err = msg.Write(output.ID()); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetTitle is for set surface title
//
// Set a short title for the surface.
//
// This string may be used to identify the surface in a task bar,
// window list, or other user interface elements provided by the
// compositor.
//
// The string must be encoded in UTF-8.
func (o *WlShellSurface) SetTitle(title wire.String) error {
	msg, err := wire.NewMessage(o.ID(), 8)
	if err != nil {
		return err
	}

	if err = msg.Write(title); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetClass is for set surface class
//
// Set a class for the surface.
//
// The surface class identifies the general class of applications
// to which the surface belongs. A common convention is to use the
// file name (or the full path if it is a non-standard location) of
// the application's .desktop file as the class.
func (o *WlShellSurface) SetClass(class wire.String) error {
	msg, err := wire.NewMessage(o.ID(), 9)
	if err != nil {
		return err
	}

	if err = msg.Write(class); err != nil {
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

// ServePing is for ping client.
//
// Ping a client to check if it is receiving events and sending
// requests. A client is expected to reply with a pong request.
func (o *WlShellSurface) ServePing(serial wire.Uint) error {
	if o.pingHandler == nil {
		return nil
	}
	return o.pingHandler(serial)
}

// HandlePing registers a handler for a Ping event.
func (o *WlShellSurface) HandlePing(h WlShellSurfacePingHandler) {
	o.pingHandler = h
}

// ServeConfigure is for suggest resize.
//
// The configure event asks the client to resize its surface.
//
// The size is a hint, in the sense that the client is free to
// ignore it if it doesn't resize, pick a smaller size (to
// satisfy aspect ratio or resize in steps of NxM pixels).
//
// The edges parameter provides a hint about how the surface
// was resized. The client may use this information to decide
// how to adjust its content to the new size (e.g. a scrolling
// area might adjust its content position to leave the viewable
// content unmoved).
//
// The client is free to dismiss all but the last configure
// event it received.
//
// The width and height arguments specify the size of the window
// in surface-local coordinates.
func (o *WlShellSurface) ServeConfigure(edges wire.Uint, width wire.Int, height wire.Int) error {
	if o.configureHandler == nil {
		return nil
	}
	return o.configureHandler(edges, width, height)
}

// HandleConfigure registers a handler for a Configure event.
func (o *WlShellSurface) HandleConfigure(h WlShellSurfaceConfigureHandler) {
	o.configureHandler = h
}

// ServePopupDone is for popup interaction is done.
//
// The popup_done event is sent out when a popup grab is broken,
// that is, when the user clicks a surface that doesn't belong
// to the client owning the popup surface.
func (o *WlShellSurface) ServePopupDone() error {
	if o.popupDoneHandler == nil {
		return nil
	}
	return o.popupDoneHandler()
}

// HandlePopupDone registers a handler for a PopupDone event.
func (o *WlShellSurface) HandlePopupDone(h WlShellSurfacePopupDoneHandler) {
	o.popupDoneHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlShellSurface) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.pingHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		serial, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.pingHandler(serial)

	case 1:
		if o.configureHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		edges, err := r.ReadUint()
		if err != nil {
			return err
		}

		width, err := r.ReadInt()
		if err != nil {
			return err
		}

		height, err := r.ReadInt()
		if err != nil {
			return err
		}

		return o.configureHandler(edges, width, height)

	case 2:
		if o.popupDoneHandler == nil {
			return nil
		}

		return o.popupDoneHandler()

	default:
		return fmt.Errorf("WlShellSurface: unhandled message(%v)", msg)
	}
}
