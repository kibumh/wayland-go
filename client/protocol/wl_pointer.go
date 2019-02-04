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

// WlPointerError is for .
//
//
type WlPointerError int

const (
	WlPointerErrorRole WlPointerError = 0 // given wl_surface has another role
)

// WlPointerButtonState is for physical button state.
//
// Describes the physical state of a button that produced the button
// event.
type WlPointerButtonState int

const (
	WlPointerButtonStateReleased WlPointerButtonState = 0 // the button is not pressed
	WlPointerButtonStatePressed  WlPointerButtonState = 1 // the button is pressed
)

// WlPointerAxis is for axis types.
//
// Describes the axis types of scroll events.
type WlPointerAxis int

const (
	WlPointerAxisVerticalScroll   WlPointerAxis = 0 // vertical axis
	WlPointerAxisHorizontalScroll WlPointerAxis = 1 // horizontal axis
)

// WlPointerAxisSource is for axis source types.
//
// Describes the source types for axis events. This indicates to the
// client how an axis event was physically generated; a client may
// adjust the user interface accordingly. For example, scroll events
// from a "finger" source may be in a smooth coordinate space with
// kinetic scrolling whereas a "wheel" source may be in discrete steps
// of a number of lines.
//
// The "continuous" axis source is a device generating events in a
// continuous coordinate space, but using something other than a
// finger. One example for this source is button-based scrolling where
// the vertical motion of a device is converted to scroll events while
// a button is held down.
//
// The "wheel tilt" axis source indicates that the actual device is a
// wheel but the scroll event is not caused by a rotation but a
// (usually sideways) tilt of the wheel.
type WlPointerAxisSource int

const (
	WlPointerAxisSourceWheel      WlPointerAxisSource = 0 // a physical wheel rotation
	WlPointerAxisSourceFinger     WlPointerAxisSource = 1 // finger on a touch surface
	WlPointerAxisSourceContinuous WlPointerAxisSource = 2 // continuous coordinate space
	WlPointerAxisSourceWheelTilt  WlPointerAxisSource = 3 // a physical wheel tilt
)

/*
 * EVENT HANDLER TYPES
 */

// WlPointerEnterHandler is a handler for enter event.
//
// Notification that this seat's pointer is focused on a certain
// surface.
//
// When a seat's focus enters a surface, the pointer image
// is undefined and a client should respond to this event by setting
// an appropriate pointer image with the set_cursor request.
type WlPointerEnterHandler func(serial wire.Uint, surface *WlSurface, surfaceX wire.Fixed, surfaceY wire.Fixed) error

// WlPointerLeaveHandler is a handler for leave event.
//
// Notification that this seat's pointer is no longer focused on
// a certain surface.
//
// The leave notification is sent before the enter notification
// for the new focus.
type WlPointerLeaveHandler func(serial wire.Uint, surface *WlSurface) error

// WlPointerMotionHandler is a handler for pointer motion event.
//
// Notification of pointer location change. The arguments
// surface_x and surface_y are the location relative to the
// focused surface.
type WlPointerMotionHandler func(time wire.Uint, surfaceX wire.Fixed, surfaceY wire.Fixed) error

// WlPointerButtonHandler is a handler for pointer button event.
//
// Mouse button click and release notifications.
//
// The location of the click is given by the last motion or
// enter event.
// The time argument is a timestamp with millisecond
// granularity, with an undefined base.
//
// The button is a button code as defined in the Linux kernel's
// linux/input-event-codes.h header file, e.g. BTN_LEFT.
//
// Any 16-bit button code value is reserved for future additions to the
// kernel's event code list. All other button codes above 0xFFFF are
// currently undefined but may be used in future versions of this
// protocol.
type WlPointerButtonHandler func(serial wire.Uint, time wire.Uint, button wire.Uint, state wire.Uint) error

// WlPointerAxisHandler is a handler for axis event.
//
// Scroll and other axis notifications.
//
// For scroll events (vertical and horizontal scroll axes), the
// value parameter is the length of a vector along the specified
// axis in a coordinate space identical to those of motion events,
// representing a relative movement along the specified axis.
//
// For devices that support movements non-parallel to axes multiple
// axis events will be emitted.
//
// When applicable, for example for touch pads, the server can
// choose to emit scroll events where the motion vector is
// equivalent to a motion event vector.
//
// When applicable, a client can transform its content relative to the
// scroll distance.
type WlPointerAxisHandler func(time wire.Uint, axis wire.Uint, value wire.Fixed) error

// WlPointerFrameHandler is a handler for end of a pointer event sequence.
//
// Indicates the end of a set of events that logically belong together.
// A client is expected to accumulate the data in all events within the
// frame before proceeding.
//
// All wl_pointer events before a wl_pointer.frame event belong
// logically together. For example, in a diagonal scroll motion the
// compositor will send an optional wl_pointer.axis_source event, two
// wl_pointer.axis events (horizontal and vertical) and finally a
// wl_pointer.frame event. The client may use this information to
// calculate a diagonal vector for scrolling.
//
// When multiple wl_pointer.axis events occur within the same frame,
// the motion vector is the combined motion of all events.
// When a wl_pointer.axis and a wl_pointer.axis_stop event occur within
// the same frame, this indicates that axis movement in one axis has
// stopped but continues in the other axis.
// When multiple wl_pointer.axis_stop events occur within the same
// frame, this indicates that these axes stopped in the same instance.
//
// A wl_pointer.frame event is sent for every logical event group,
// even if the group only contains a single wl_pointer event.
// Specifically, a client may get a sequence: motion, frame, button,
// frame, axis, frame, axis_stop, frame.
//
// The wl_pointer.enter and wl_pointer.leave events are logical events
// generated by the compositor and not the hardware. These events are
// also grouped by a wl_pointer.frame. When a pointer moves from one
// surface to another, a compositor should group the
// wl_pointer.leave event within the same wl_pointer.frame.
// However, a client must not rely on wl_pointer.leave and
// wl_pointer.enter being in the same wl_pointer.frame.
// Compositor-specific policies may require the wl_pointer.leave and
// wl_pointer.enter event being split across multiple wl_pointer.frame
// groups.
type WlPointerFrameHandler func() error

// WlPointerAxisSourceHandler is a handler for axis source event.
//
// Source information for scroll and other axes.
//
// This event does not occur on its own. It is sent before a
// wl_pointer.frame event and carries the source information for
// all events within that frame.
//
// The source specifies how this event was generated. If the source is
// wl_pointer.axis_source.finger, a wl_pointer.axis_stop event will be
// sent when the user lifts the finger off the device.
//
// If the source is wl_pointer.axis_source.wheel,
// wl_pointer.axis_source.wheel_tilt or
// wl_pointer.axis_source.continuous, a wl_pointer.axis_stop event may
// or may not be sent. Whether a compositor sends an axis_stop event
// for these sources is hardware-specific and implementation-dependent;
// clients must not rely on receiving an axis_stop event for these
// scroll sources and should treat scroll sequences from these scroll
// sources as unterminated by default.
//
// This event is optional. If the source is unknown for a particular
// axis event sequence, no event is sent.
// Only one wl_pointer.axis_source event is permitted per frame.
//
// The order of wl_pointer.axis_discrete and wl_pointer.axis_source is
// not guaranteed.
type WlPointerAxisSourceHandler func(axisSource wire.Uint) error

// WlPointerAxisStopHandler is a handler for axis stop event.
//
// Stop notification for scroll and other axes.
//
// For some wl_pointer.axis_source types, a wl_pointer.axis_stop event
// is sent to notify a client that the axis sequence has terminated.
// This enables the client to implement kinetic scrolling.
// See the wl_pointer.axis_source documentation for information on when
// this event may be generated.
//
// Any wl_pointer.axis events with the same axis_source after this
// event should be considered as the start of a new axis motion.
//
// The timestamp is to be interpreted identical to the timestamp in the
// wl_pointer.axis event. The timestamp value may be the same as a
// preceding wl_pointer.axis event.
type WlPointerAxisStopHandler func(time wire.Uint, axis wire.Uint) error

// WlPointerAxisDiscreteHandler is a handler for axis click event.
//
// Discrete step information for scroll and other axes.
//
// This event carries the axis value of the wl_pointer.axis event in
// discrete steps (e.g. mouse wheel clicks).
//
// This event does not occur on its own, it is coupled with a
// wl_pointer.axis event that represents this axis value on a
// continuous scale. The protocol guarantees that each axis_discrete
// event is always followed by exactly one axis event with the same
// axis number within the same wl_pointer.frame. Note that the protocol
// allows for other events to occur between the axis_discrete and
// its coupled axis event, including other axis_discrete or axis
// events.
//
// This event is optional; continuous scrolling devices
// like two-finger scrolling on touchpads do not have discrete
// steps and do not generate this event.
//
// The discrete value carries the directional information. e.g. a value
// of -2 is two steps towards the negative direction of this axis.
//
// The axis number is identical to the axis number in the associated
// axis event.
//
// The order of wl_pointer.axis_discrete and wl_pointer.axis_source is
// not guaranteed.
type WlPointerAxisDiscreteHandler func(axis wire.Uint, discrete wire.Int) error

/*
 * TYPE
 */
// WlPointer is pointer input device.
//
// The wl_pointer interface represents one or more input devices,
// such as mice, which control the pointer location and pointer_focus
// of a seat.
//
// The wl_pointer interface generates motion, enter and leave
// events for the surfaces that the pointer is located over,
// and button and axis events for button presses, button releases
// and scrolling.
type WlPointer struct {
	Base

	enterHandler        WlPointerEnterHandler
	leaveHandler        WlPointerLeaveHandler
	motionHandler       WlPointerMotionHandler
	buttonHandler       WlPointerButtonHandler
	axisHandler         WlPointerAxisHandler
	frameHandler        WlPointerFrameHandler
	axisSourceHandler   WlPointerAxisSourceHandler
	axisStopHandler     WlPointerAxisStopHandler
	axisDiscreteHandler WlPointerAxisDiscreteHandler
}

// NewWlPointer creates a WlPointer object.
func NewWlPointer(c *wire.Conn) *WlPointer {
	return NewWlPointerWithID(c, c.NewID())
}

// NewWlPointerWithID creates a WlPointer object with a given id.
func NewWlPointerWithID(c *wire.Conn, id wire.ID) *WlPointer {
	o := &WlPointer{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// SetCursor is for set the pointer surface
//
// Set the pointer surface, i.e., the surface that contains the
// pointer image (cursor). This request gives the surface the role
// of a cursor. If the surface already has another role, it raises
// a protocol error.
//
// The cursor actually changes only if the pointer
// focus for this device is one of the requesting client's surfaces
// or the surface parameter is the current pointer surface. If
// there was a previous surface set with this request it is
// replaced. If surface is NULL, the pointer image is hidden.
//
// The parameters hotspot_x and hotspot_y define the position of
// the pointer surface relative to the pointer location. Its
// top-left corner is always at (x, y) - (hotspot_x, hotspot_y),
// where (x, y) are the coordinates of the pointer location, in
// surface-local coordinates.
//
// On surface.attach requests to the pointer surface, hotspot_x
// and hotspot_y are decremented by the x and y parameters
// passed to the request. Attach must be confirmed by
// wl_surface.commit as usual.
//
// The hotspot can also be updated by passing the currently set
// pointer surface to this request with new values for hotspot_x
// and hotspot_y.
//
// The current and pending input regions of the wl_surface are
// cleared, and wl_surface.set_input_region is ignored until the
// wl_surface is no longer used as the cursor. When the use as a
// cursor ends, the current and pending input regions become
// undefined, and the wl_surface is unmapped.
func (o *WlPointer) SetCursor(serial wire.Uint, surface *WlSurface, hotspotX wire.Int, hotspotY wire.Int) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(serial); err != nil {
		return err
	}

	if err = msg.Write(surface.ID()); err != nil {
		return err
	}

	if err = msg.Write(hotspotX); err != nil {
		return err
	}

	if err = msg.Write(hotspotY); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Release is for release the pointer object
//
// Using this request a client can tell the server that it is not going to
// use the pointer object anymore.
//
// This request destroys the pointer proxy object, so clients must not call
// wl_pointer_destroy() after using this request.
func (o *WlPointer) Release() error {
	msg, err := wire.NewMessage(o.ID(), 1)
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

// ServeEnter is for enter event.
//
// Notification that this seat's pointer is focused on a certain
// surface.
//
// When a seat's focus enters a surface, the pointer image
// is undefined and a client should respond to this event by setting
// an appropriate pointer image with the set_cursor request.
func (o *WlPointer) ServeEnter(serial wire.Uint, surface *WlSurface, surfaceX wire.Fixed, surfaceY wire.Fixed) error {
	if o.enterHandler == nil {
		return nil
	}
	return o.enterHandler(serial, surface, surfaceX, surfaceY)
}

// HandleEnter registers a handler for a Enter event.
func (o *WlPointer) HandleEnter(h WlPointerEnterHandler) {
	o.enterHandler = h
}

// ServeLeave is for leave event.
//
// Notification that this seat's pointer is no longer focused on
// a certain surface.
//
// The leave notification is sent before the enter notification
// for the new focus.
func (o *WlPointer) ServeLeave(serial wire.Uint, surface *WlSurface) error {
	if o.leaveHandler == nil {
		return nil
	}
	return o.leaveHandler(serial, surface)
}

// HandleLeave registers a handler for a Leave event.
func (o *WlPointer) HandleLeave(h WlPointerLeaveHandler) {
	o.leaveHandler = h
}

// ServeMotion is for pointer motion event.
//
// Notification of pointer location change. The arguments
// surface_x and surface_y are the location relative to the
// focused surface.
func (o *WlPointer) ServeMotion(time wire.Uint, surfaceX wire.Fixed, surfaceY wire.Fixed) error {
	if o.motionHandler == nil {
		return nil
	}
	return o.motionHandler(time, surfaceX, surfaceY)
}

// HandleMotion registers a handler for a Motion event.
func (o *WlPointer) HandleMotion(h WlPointerMotionHandler) {
	o.motionHandler = h
}

// ServeButton is for pointer button event.
//
// Mouse button click and release notifications.
//
// The location of the click is given by the last motion or
// enter event.
// The time argument is a timestamp with millisecond
// granularity, with an undefined base.
//
// The button is a button code as defined in the Linux kernel's
// linux/input-event-codes.h header file, e.g. BTN_LEFT.
//
// Any 16-bit button code value is reserved for future additions to the
// kernel's event code list. All other button codes above 0xFFFF are
// currently undefined but may be used in future versions of this
// protocol.
func (o *WlPointer) ServeButton(serial wire.Uint, time wire.Uint, button wire.Uint, state wire.Uint) error {
	if o.buttonHandler == nil {
		return nil
	}
	return o.buttonHandler(serial, time, button, state)
}

// HandleButton registers a handler for a Button event.
func (o *WlPointer) HandleButton(h WlPointerButtonHandler) {
	o.buttonHandler = h
}

// ServeAxis is for axis event.
//
// Scroll and other axis notifications.
//
// For scroll events (vertical and horizontal scroll axes), the
// value parameter is the length of a vector along the specified
// axis in a coordinate space identical to those of motion events,
// representing a relative movement along the specified axis.
//
// For devices that support movements non-parallel to axes multiple
// axis events will be emitted.
//
// When applicable, for example for touch pads, the server can
// choose to emit scroll events where the motion vector is
// equivalent to a motion event vector.
//
// When applicable, a client can transform its content relative to the
// scroll distance.
func (o *WlPointer) ServeAxis(time wire.Uint, axis wire.Uint, value wire.Fixed) error {
	if o.axisHandler == nil {
		return nil
	}
	return o.axisHandler(time, axis, value)
}

// HandleAxis registers a handler for a Axis event.
func (o *WlPointer) HandleAxis(h WlPointerAxisHandler) {
	o.axisHandler = h
}

// ServeFrame is for end of a pointer event sequence.
//
// Indicates the end of a set of events that logically belong together.
// A client is expected to accumulate the data in all events within the
// frame before proceeding.
//
// All wl_pointer events before a wl_pointer.frame event belong
// logically together. For example, in a diagonal scroll motion the
// compositor will send an optional wl_pointer.axis_source event, two
// wl_pointer.axis events (horizontal and vertical) and finally a
// wl_pointer.frame event. The client may use this information to
// calculate a diagonal vector for scrolling.
//
// When multiple wl_pointer.axis events occur within the same frame,
// the motion vector is the combined motion of all events.
// When a wl_pointer.axis and a wl_pointer.axis_stop event occur within
// the same frame, this indicates that axis movement in one axis has
// stopped but continues in the other axis.
// When multiple wl_pointer.axis_stop events occur within the same
// frame, this indicates that these axes stopped in the same instance.
//
// A wl_pointer.frame event is sent for every logical event group,
// even if the group only contains a single wl_pointer event.
// Specifically, a client may get a sequence: motion, frame, button,
// frame, axis, frame, axis_stop, frame.
//
// The wl_pointer.enter and wl_pointer.leave events are logical events
// generated by the compositor and not the hardware. These events are
// also grouped by a wl_pointer.frame. When a pointer moves from one
// surface to another, a compositor should group the
// wl_pointer.leave event within the same wl_pointer.frame.
// However, a client must not rely on wl_pointer.leave and
// wl_pointer.enter being in the same wl_pointer.frame.
// Compositor-specific policies may require the wl_pointer.leave and
// wl_pointer.enter event being split across multiple wl_pointer.frame
// groups.
func (o *WlPointer) ServeFrame() error {
	if o.frameHandler == nil {
		return nil
	}
	return o.frameHandler()
}

// HandleFrame registers a handler for a Frame event.
func (o *WlPointer) HandleFrame(h WlPointerFrameHandler) {
	o.frameHandler = h
}

// ServeAxisSource is for axis source event.
//
// Source information for scroll and other axes.
//
// This event does not occur on its own. It is sent before a
// wl_pointer.frame event and carries the source information for
// all events within that frame.
//
// The source specifies how this event was generated. If the source is
// wl_pointer.axis_source.finger, a wl_pointer.axis_stop event will be
// sent when the user lifts the finger off the device.
//
// If the source is wl_pointer.axis_source.wheel,
// wl_pointer.axis_source.wheel_tilt or
// wl_pointer.axis_source.continuous, a wl_pointer.axis_stop event may
// or may not be sent. Whether a compositor sends an axis_stop event
// for these sources is hardware-specific and implementation-dependent;
// clients must not rely on receiving an axis_stop event for these
// scroll sources and should treat scroll sequences from these scroll
// sources as unterminated by default.
//
// This event is optional. If the source is unknown for a particular
// axis event sequence, no event is sent.
// Only one wl_pointer.axis_source event is permitted per frame.
//
// The order of wl_pointer.axis_discrete and wl_pointer.axis_source is
// not guaranteed.
func (o *WlPointer) ServeAxisSource(axisSource wire.Uint) error {
	if o.axisSourceHandler == nil {
		return nil
	}
	return o.axisSourceHandler(axisSource)
}

// HandleAxisSource registers a handler for a AxisSource event.
func (o *WlPointer) HandleAxisSource(h WlPointerAxisSourceHandler) {
	o.axisSourceHandler = h
}

// ServeAxisStop is for axis stop event.
//
// Stop notification for scroll and other axes.
//
// For some wl_pointer.axis_source types, a wl_pointer.axis_stop event
// is sent to notify a client that the axis sequence has terminated.
// This enables the client to implement kinetic scrolling.
// See the wl_pointer.axis_source documentation for information on when
// this event may be generated.
//
// Any wl_pointer.axis events with the same axis_source after this
// event should be considered as the start of a new axis motion.
//
// The timestamp is to be interpreted identical to the timestamp in the
// wl_pointer.axis event. The timestamp value may be the same as a
// preceding wl_pointer.axis event.
func (o *WlPointer) ServeAxisStop(time wire.Uint, axis wire.Uint) error {
	if o.axisStopHandler == nil {
		return nil
	}
	return o.axisStopHandler(time, axis)
}

// HandleAxisStop registers a handler for a AxisStop event.
func (o *WlPointer) HandleAxisStop(h WlPointerAxisStopHandler) {
	o.axisStopHandler = h
}

// ServeAxisDiscrete is for axis click event.
//
// Discrete step information for scroll and other axes.
//
// This event carries the axis value of the wl_pointer.axis event in
// discrete steps (e.g. mouse wheel clicks).
//
// This event does not occur on its own, it is coupled with a
// wl_pointer.axis event that represents this axis value on a
// continuous scale. The protocol guarantees that each axis_discrete
// event is always followed by exactly one axis event with the same
// axis number within the same wl_pointer.frame. Note that the protocol
// allows for other events to occur between the axis_discrete and
// its coupled axis event, including other axis_discrete or axis
// events.
//
// This event is optional; continuous scrolling devices
// like two-finger scrolling on touchpads do not have discrete
// steps and do not generate this event.
//
// The discrete value carries the directional information. e.g. a value
// of -2 is two steps towards the negative direction of this axis.
//
// The axis number is identical to the axis number in the associated
// axis event.
//
// The order of wl_pointer.axis_discrete and wl_pointer.axis_source is
// not guaranteed.
func (o *WlPointer) ServeAxisDiscrete(axis wire.Uint, discrete wire.Int) error {
	if o.axisDiscreteHandler == nil {
		return nil
	}
	return o.axisDiscreteHandler(axis, discrete)
}

// HandleAxisDiscrete registers a handler for a AxisDiscrete event.
func (o *WlPointer) HandleAxisDiscrete(h WlPointerAxisDiscreteHandler) {
	o.axisDiscreteHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlPointer) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.enterHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		serial, err := r.ReadUint()
		if err != nil {
			return err
		}

		surfaceId, err := r.ReadID()
		if err != nil {
			return err
		}

		surfaceObj, ok := o.Base.Conn.GetObject(surfaceId)
		if !ok {
			return fmt.Errorf("cannot find an object: id(%d)", surfaceId)
		}
		surface, ok := surfaceObj.(*WlSurface)
		if !ok {
			return fmt.Errorf("failed to type assertion: id(%d), type(WlSurface)", surfaceId)
		}

		surfaceX, err := r.ReadFixed()
		if err != nil {
			return err
		}

		surfaceY, err := r.ReadFixed()
		if err != nil {
			return err
		}

		return o.enterHandler(serial, surface, surfaceX, surfaceY)

	case 1:
		if o.leaveHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		serial, err := r.ReadUint()
		if err != nil {
			return err
		}

		surfaceId, err := r.ReadID()
		if err != nil {
			return err
		}

		surfaceObj, ok := o.Base.Conn.GetObject(surfaceId)
		if !ok {
			return fmt.Errorf("cannot find an object: id(%d)", surfaceId)
		}
		surface, ok := surfaceObj.(*WlSurface)
		if !ok {
			return fmt.Errorf("failed to type assertion: id(%d), type(WlSurface)", surfaceId)
		}

		return o.leaveHandler(serial, surface)

	case 2:
		if o.motionHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		time, err := r.ReadUint()
		if err != nil {
			return err
		}

		surfaceX, err := r.ReadFixed()
		if err != nil {
			return err
		}

		surfaceY, err := r.ReadFixed()
		if err != nil {
			return err
		}

		return o.motionHandler(time, surfaceX, surfaceY)

	case 3:
		if o.buttonHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		serial, err := r.ReadUint()
		if err != nil {
			return err
		}

		time, err := r.ReadUint()
		if err != nil {
			return err
		}

		button, err := r.ReadUint()
		if err != nil {
			return err
		}

		state, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.buttonHandler(serial, time, button, state)

	case 4:
		if o.axisHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		time, err := r.ReadUint()
		if err != nil {
			return err
		}

		axis, err := r.ReadUint()
		if err != nil {
			return err
		}

		value, err := r.ReadFixed()
		if err != nil {
			return err
		}

		return o.axisHandler(time, axis, value)

	case 5:
		if o.frameHandler == nil {
			return nil
		}

		return o.frameHandler()

	case 6:
		if o.axisSourceHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		axisSource, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.axisSourceHandler(axisSource)

	case 7:
		if o.axisStopHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		time, err := r.ReadUint()
		if err != nil {
			return err
		}

		axis, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.axisStopHandler(time, axis)

	case 8:
		if o.axisDiscreteHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		axis, err := r.ReadUint()
		if err != nil {
			return err
		}

		discrete, err := r.ReadInt()
		if err != nil {
			return err
		}

		return o.axisDiscreteHandler(axis, discrete)

	default:
		return fmt.Errorf("WlPointer: unhandled message(%v)", msg)
	}
}
