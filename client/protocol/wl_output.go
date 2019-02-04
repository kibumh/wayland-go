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

// WlOutputSubpixel is for subpixel geometry information.
//
// This enumeration describes how the physical
// pixels on an output are laid out.
type WlOutputSubpixel int

const (
	WlOutputSubpixelUnknown       WlOutputSubpixel = 0 // unknown geometry
	WlOutputSubpixelNone          WlOutputSubpixel = 1 // no geometry
	WlOutputSubpixelHorizontalRgb WlOutputSubpixel = 2 // horizontal RGB
	WlOutputSubpixelHorizontalBgr WlOutputSubpixel = 3 // horizontal BGR
	WlOutputSubpixelVerticalRgb   WlOutputSubpixel = 4 // vertical RGB
	WlOutputSubpixelVerticalBgr   WlOutputSubpixel = 5 // vertical BGR
)

// WlOutputTransform is for transform from framebuffer to output.
//
// This describes the transform that a compositor will apply to a
// surface to compensate for the rotation or mirroring of an
// output device.
//
// The flipped values correspond to an initial flip around a
// vertical axis followed by rotation.
//
// The purpose is mainly to allow clients to render accordingly and
// tell the compositor, so that for fullscreen surfaces, the
// compositor will still be able to scan out directly from client
// surfaces.
type WlOutputTransform int

const (
	WlOutputTransformNormal     WlOutputTransform = 0 // no transform
	WlOutputTransform90         WlOutputTransform = 1 // 90 degrees counter-clockwise
	WlOutputTransform180        WlOutputTransform = 2 // 180 degrees counter-clockwise
	WlOutputTransform270        WlOutputTransform = 3 // 270 degrees counter-clockwise
	WlOutputTransformFlipped    WlOutputTransform = 4 // 180 degree flip around a vertical axis
	WlOutputTransformFlipped90  WlOutputTransform = 5 // flip and rotate 90 degrees counter-clockwise
	WlOutputTransformFlipped180 WlOutputTransform = 6 // flip and rotate 180 degrees counter-clockwise
	WlOutputTransformFlipped270 WlOutputTransform = 7 // flip and rotate 270 degrees counter-clockwise
)

// WlOutputMode is for mode information.
//
// These flags describe properties of an output mode.
// They are used in the flags bitfield of the mode event.
type WlOutputMode int

const (
	WlOutputModeCurrent   WlOutputMode = 0x1 // indicates this is the current mode
	WlOutputModePreferred WlOutputMode = 0x2 // indicates this is the preferred mode
)

/*
 * EVENT HANDLER TYPES
 */

// WlOutputGeometryHandler is a handler for properties of the output.
//
// The geometry event describes geometric properties of the output.
// The event is sent when binding to the output object and whenever
// any of the properties change.
//
// The physical size can be set to zero if it doesn't make sense for this
// output (e.g. for projectors or virtual outputs).
type WlOutputGeometryHandler func(x wire.Int, y wire.Int, physicalWidth wire.Int, physicalHeight wire.Int, subpixel wire.Int, make wire.String, model wire.String, transform wire.Int) error

// WlOutputModeHandler is a handler for advertise available modes for the output.
//
// The mode event describes an available mode for the output.
//
// The event is sent when binding to the output object and there
// will always be one mode, the current mode.  The event is sent
// again if an output changes mode, for the mode that is now
// current.  In other words, the current mode is always the last
// mode that was received with the current flag set.
//
// The size of a mode is given in physical hardware units of
// the output device. This is not necessarily the same as
// the output size in the global compositor space. For instance,
// the output may be scaled, as described in wl_output.scale,
// or transformed, as described in wl_output.transform.
type WlOutputModeHandler func(flags wire.Uint, width wire.Int, height wire.Int, refresh wire.Int) error

// WlOutputDoneHandler is a handler for sent all information about output.
//
// This event is sent after all other properties have been
// sent after binding to the output object and after any
// other property changes done after that. This allows
// changes to the output properties to be seen as
// atomic, even if they happen via multiple events.
type WlOutputDoneHandler func() error

// WlOutputScaleHandler is a handler for output scaling properties.
//
// This event contains scaling geometry information
// that is not in the geometry event. It may be sent after
// binding the output object or if the output scale changes
// later. If it is not sent, the client should assume a
// scale of 1.
//
// A scale larger than 1 means that the compositor will
// automatically scale surface buffers by this amount
// when rendering. This is used for very high resolution
// displays where applications rendering at the native
// resolution would be too small to be legible.
//
// It is intended that scaling aware clients track the
// current output of a surface, and if it is on a scaled
// output it should use wl_surface.set_buffer_scale with
// the scale of the output. That way the compositor can
// avoid scaling the surface, and the client can supply
// a higher detail image.
type WlOutputScaleHandler func(factor wire.Int) error

/*
 * TYPE
 */
// WlOutput is compositor output region.
//
// An output describes part of the compositor geometry.  The
// compositor works in the 'compositor coordinate system' and an
// output corresponds to a rectangular area in that space that is
// actually visible.  This typically corresponds to a monitor that
// displays part of the compositor space.  This object is published
// as global during start up, or when a monitor is hotplugged.
type WlOutput struct {
	Base

	geometryHandler WlOutputGeometryHandler
	modeHandler     WlOutputModeHandler
	doneHandler     WlOutputDoneHandler
	scaleHandler    WlOutputScaleHandler
}

// NewWlOutput creates a WlOutput object.
func NewWlOutput(c *wire.Conn) *WlOutput {
	return NewWlOutputWithID(c, c.NewID())
}

// NewWlOutputWithID creates a WlOutput object with a given id.
func NewWlOutputWithID(c *wire.Conn, id wire.ID) *WlOutput {
	o := &WlOutput{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Release is for release the output object
//
// Using this request a client can tell the server that it is not going to
// use the output object anymore.
func (o *WlOutput) Release() error {
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

// ServeGeometry is for properties of the output.
//
// The geometry event describes geometric properties of the output.
// The event is sent when binding to the output object and whenever
// any of the properties change.
//
// The physical size can be set to zero if it doesn't make sense for this
// output (e.g. for projectors or virtual outputs).
func (o *WlOutput) ServeGeometry(x wire.Int, y wire.Int, physicalWidth wire.Int, physicalHeight wire.Int, subpixel wire.Int, make wire.String, model wire.String, transform wire.Int) error {
	if o.geometryHandler == nil {
		return nil
	}
	return o.geometryHandler(x, y, physicalWidth, physicalHeight, subpixel, make, model, transform)
}

// HandleGeometry registers a handler for a Geometry event.
func (o *WlOutput) HandleGeometry(h WlOutputGeometryHandler) {
	o.geometryHandler = h
}

// ServeMode is for advertise available modes for the output.
//
// The mode event describes an available mode for the output.
//
// The event is sent when binding to the output object and there
// will always be one mode, the current mode.  The event is sent
// again if an output changes mode, for the mode that is now
// current.  In other words, the current mode is always the last
// mode that was received with the current flag set.
//
// The size of a mode is given in physical hardware units of
// the output device. This is not necessarily the same as
// the output size in the global compositor space. For instance,
// the output may be scaled, as described in wl_output.scale,
// or transformed, as described in wl_output.transform.
func (o *WlOutput) ServeMode(flags wire.Uint, width wire.Int, height wire.Int, refresh wire.Int) error {
	if o.modeHandler == nil {
		return nil
	}
	return o.modeHandler(flags, width, height, refresh)
}

// HandleMode registers a handler for a Mode event.
func (o *WlOutput) HandleMode(h WlOutputModeHandler) {
	o.modeHandler = h
}

// ServeDone is for sent all information about output.
//
// This event is sent after all other properties have been
// sent after binding to the output object and after any
// other property changes done after that. This allows
// changes to the output properties to be seen as
// atomic, even if they happen via multiple events.
func (o *WlOutput) ServeDone() error {
	if o.doneHandler == nil {
		return nil
	}
	return o.doneHandler()
}

// HandleDone registers a handler for a Done event.
func (o *WlOutput) HandleDone(h WlOutputDoneHandler) {
	o.doneHandler = h
}

// ServeScale is for output scaling properties.
//
// This event contains scaling geometry information
// that is not in the geometry event. It may be sent after
// binding the output object or if the output scale changes
// later. If it is not sent, the client should assume a
// scale of 1.
//
// A scale larger than 1 means that the compositor will
// automatically scale surface buffers by this amount
// when rendering. This is used for very high resolution
// displays where applications rendering at the native
// resolution would be too small to be legible.
//
// It is intended that scaling aware clients track the
// current output of a surface, and if it is on a scaled
// output it should use wl_surface.set_buffer_scale with
// the scale of the output. That way the compositor can
// avoid scaling the surface, and the client can supply
// a higher detail image.
func (o *WlOutput) ServeScale(factor wire.Int) error {
	if o.scaleHandler == nil {
		return nil
	}
	return o.scaleHandler(factor)
}

// HandleScale registers a handler for a Scale event.
func (o *WlOutput) HandleScale(h WlOutputScaleHandler) {
	o.scaleHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlOutput) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.geometryHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		x, err := r.ReadInt()
		if err != nil {
			return err
		}

		y, err := r.ReadInt()
		if err != nil {
			return err
		}

		physicalWidth, err := r.ReadInt()
		if err != nil {
			return err
		}

		physicalHeight, err := r.ReadInt()
		if err != nil {
			return err
		}

		subpixel, err := r.ReadInt()
		if err != nil {
			return err
		}

		make, err := r.ReadString()
		if err != nil {
			return err
		}

		model, err := r.ReadString()
		if err != nil {
			return err
		}

		transform, err := r.ReadInt()
		if err != nil {
			return err
		}

		return o.geometryHandler(x, y, physicalWidth, physicalHeight, subpixel, make, model, transform)

	case 1:
		if o.modeHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		flags, err := r.ReadUint()
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

		refresh, err := r.ReadInt()
		if err != nil {
			return err
		}

		return o.modeHandler(flags, width, height, refresh)

	case 2:
		if o.doneHandler == nil {
			return nil
		}

		return o.doneHandler()

	case 3:
		if o.scaleHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		factor, err := r.ReadInt()
		if err != nil {
			return err
		}

		return o.scaleHandler(factor)

	default:
		return fmt.Errorf("WlOutput: unhandled message(%v)", msg)
	}
}
