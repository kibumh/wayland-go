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

// WlTouchDownHandler is a handler for touch down event and beginning of a touch sequence.
//
// A new touch point has appeared on the surface. This touch point is
// assigned a unique ID. Future events from this touch point reference
// this ID. The ID ceases to be valid after a touch up event and may be
// reused in the future.
type WlTouchDownHandler func(serial wire.Uint, time wire.Uint, surface *WlSurface, id wire.Int, x wire.Fixed, y wire.Fixed) error

// WlTouchUpHandler is a handler for end of a touch event sequence.
//
// The touch point has disappeared. No further events will be sent for
// this touch point and the touch point's ID is released and may be
// reused in a future touch down event.
type WlTouchUpHandler func(serial wire.Uint, time wire.Uint, id wire.Int) error

// WlTouchMotionHandler is a handler for update of touch point coordinates.
//
// A touch point has changed coordinates.
type WlTouchMotionHandler func(time wire.Uint, id wire.Int, x wire.Fixed, y wire.Fixed) error

// WlTouchFrameHandler is a handler for end of touch frame event.
//
// Indicates the end of a set of events that logically belong together.
// A client is expected to accumulate the data in all events within the
// frame before proceeding.
//
// A wl_touch.frame terminates at least one event but otherwise no
// guarantee is provided about the set of events within a frame. A client
// must assume that any state not updated in a frame is unchanged from the
// previously known state.
type WlTouchFrameHandler func() error

// WlTouchCancelHandler is a handler for touch session cancelled.
//
// Sent if the compositor decides the touch stream is a global
// gesture. No further events are sent to the clients from that
// particular gesture. Touch cancellation applies to all touch points
// currently active on this client's surface. The client is
// responsible for finalizing the touch points, future touch points on
// this surface may reuse the touch point ID.
type WlTouchCancelHandler func() error

// WlTouchShapeHandler is a handler for update shape of touch point.
//
// Sent when a touchpoint has changed its shape.
//
// This event does not occur on its own. It is sent before a
// wl_touch.frame event and carries the new shape information for
// any previously reported, or new touch points of that frame.
//
// Other events describing the touch point such as wl_touch.down,
// wl_touch.motion or wl_touch.orientation may be sent within the
// same wl_touch.frame. A client should treat these events as a single
// logical touch point update. The order of wl_touch.shape,
// wl_touch.orientation and wl_touch.motion is not guaranteed.
// A wl_touch.down event is guaranteed to occur before the first
// wl_touch.shape event for this touch ID but both events may occur within
// the same wl_touch.frame.
//
// A touchpoint shape is approximated by an ellipse through the major and
// minor axis length. The major axis length describes the longer diameter
// of the ellipse, while the minor axis length describes the shorter
// diameter. Major and minor are orthogonal and both are specified in
// surface-local coordinates. The center of the ellipse is always at the
// touchpoint location as reported by wl_touch.down or wl_touch.move.
//
// This event is only sent by the compositor if the touch device supports
// shape reports. The client has to make reasonable assumptions about the
// shape if it did not receive this event.
type WlTouchShapeHandler func(id wire.Int, major wire.Fixed, minor wire.Fixed) error

// WlTouchOrientationHandler is a handler for update orientation of touch point.
//
// Sent when a touchpoint has changed its orientation.
//
// This event does not occur on its own. It is sent before a
// wl_touch.frame event and carries the new shape information for
// any previously reported, or new touch points of that frame.
//
// Other events describing the touch point such as wl_touch.down,
// wl_touch.motion or wl_touch.shape may be sent within the
// same wl_touch.frame. A client should treat these events as a single
// logical touch point update. The order of wl_touch.shape,
// wl_touch.orientation and wl_touch.motion is not guaranteed.
// A wl_touch.down event is guaranteed to occur before the first
// wl_touch.orientation event for this touch ID but both events may occur
// within the same wl_touch.frame.
//
// The orientation describes the clockwise angle of a touchpoint's major
// axis to the positive surface y-axis and is normalized to the -180 to
// +180 degree range. The granularity of orientation depends on the touch
// device, some devices only support binary rotation values between 0 and
// 90 degrees.
//
// This event is only sent by the compositor if the touch device supports
// orientation reports.
type WlTouchOrientationHandler func(id wire.Int, orientation wire.Fixed) error

/*
 * TYPE
 */
// WlTouch is touchscreen input device.
//
// The wl_touch interface represents a touchscreen
// associated with a seat.
//
// Touch interactions can consist of one or more contacts.
// For each contact, a series of events is generated, starting
// with a down event, followed by zero or more motion events,
// and ending with an up event. Events relating to the same
// contact point can be identified by the ID of the sequence.
type WlTouch struct {
	Base

	downHandler        WlTouchDownHandler
	upHandler          WlTouchUpHandler
	motionHandler      WlTouchMotionHandler
	frameHandler       WlTouchFrameHandler
	cancelHandler      WlTouchCancelHandler
	shapeHandler       WlTouchShapeHandler
	orientationHandler WlTouchOrientationHandler
}

// NewWlTouch creates a WlTouch object.
func NewWlTouch(c *wire.Conn) *WlTouch {
	return NewWlTouchWithID(c, c.NewID())
}

// NewWlTouchWithID creates a WlTouch object with a given id.
func NewWlTouchWithID(c *wire.Conn, id wire.ID) *WlTouch {
	o := &WlTouch{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Release is for release the touch object
//
//
func (o *WlTouch) Release() error {
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

// ServeDown is for touch down event and beginning of a touch sequence.
//
// A new touch point has appeared on the surface. This touch point is
// assigned a unique ID. Future events from this touch point reference
// this ID. The ID ceases to be valid after a touch up event and may be
// reused in the future.
func (o *WlTouch) ServeDown(serial wire.Uint, time wire.Uint, surface *WlSurface, id wire.Int, x wire.Fixed, y wire.Fixed) error {
	if o.downHandler == nil {
		return nil
	}
	return o.downHandler(serial, time, surface, id, x, y)
}

// HandleDown registers a handler for a Down event.
func (o *WlTouch) HandleDown(h WlTouchDownHandler) {
	o.downHandler = h
}

// ServeUp is for end of a touch event sequence.
//
// The touch point has disappeared. No further events will be sent for
// this touch point and the touch point's ID is released and may be
// reused in a future touch down event.
func (o *WlTouch) ServeUp(serial wire.Uint, time wire.Uint, id wire.Int) error {
	if o.upHandler == nil {
		return nil
	}
	return o.upHandler(serial, time, id)
}

// HandleUp registers a handler for a Up event.
func (o *WlTouch) HandleUp(h WlTouchUpHandler) {
	o.upHandler = h
}

// ServeMotion is for update of touch point coordinates.
//
// A touch point has changed coordinates.
func (o *WlTouch) ServeMotion(time wire.Uint, id wire.Int, x wire.Fixed, y wire.Fixed) error {
	if o.motionHandler == nil {
		return nil
	}
	return o.motionHandler(time, id, x, y)
}

// HandleMotion registers a handler for a Motion event.
func (o *WlTouch) HandleMotion(h WlTouchMotionHandler) {
	o.motionHandler = h
}

// ServeFrame is for end of touch frame event.
//
// Indicates the end of a set of events that logically belong together.
// A client is expected to accumulate the data in all events within the
// frame before proceeding.
//
// A wl_touch.frame terminates at least one event but otherwise no
// guarantee is provided about the set of events within a frame. A client
// must assume that any state not updated in a frame is unchanged from the
// previously known state.
func (o *WlTouch) ServeFrame() error {
	if o.frameHandler == nil {
		return nil
	}
	return o.frameHandler()
}

// HandleFrame registers a handler for a Frame event.
func (o *WlTouch) HandleFrame(h WlTouchFrameHandler) {
	o.frameHandler = h
}

// ServeCancel is for touch session cancelled.
//
// Sent if the compositor decides the touch stream is a global
// gesture. No further events are sent to the clients from that
// particular gesture. Touch cancellation applies to all touch points
// currently active on this client's surface. The client is
// responsible for finalizing the touch points, future touch points on
// this surface may reuse the touch point ID.
func (o *WlTouch) ServeCancel() error {
	if o.cancelHandler == nil {
		return nil
	}
	return o.cancelHandler()
}

// HandleCancel registers a handler for a Cancel event.
func (o *WlTouch) HandleCancel(h WlTouchCancelHandler) {
	o.cancelHandler = h
}

// ServeShape is for update shape of touch point.
//
// Sent when a touchpoint has changed its shape.
//
// This event does not occur on its own. It is sent before a
// wl_touch.frame event and carries the new shape information for
// any previously reported, or new touch points of that frame.
//
// Other events describing the touch point such as wl_touch.down,
// wl_touch.motion or wl_touch.orientation may be sent within the
// same wl_touch.frame. A client should treat these events as a single
// logical touch point update. The order of wl_touch.shape,
// wl_touch.orientation and wl_touch.motion is not guaranteed.
// A wl_touch.down event is guaranteed to occur before the first
// wl_touch.shape event for this touch ID but both events may occur within
// the same wl_touch.frame.
//
// A touchpoint shape is approximated by an ellipse through the major and
// minor axis length. The major axis length describes the longer diameter
// of the ellipse, while the minor axis length describes the shorter
// diameter. Major and minor are orthogonal and both are specified in
// surface-local coordinates. The center of the ellipse is always at the
// touchpoint location as reported by wl_touch.down or wl_touch.move.
//
// This event is only sent by the compositor if the touch device supports
// shape reports. The client has to make reasonable assumptions about the
// shape if it did not receive this event.
func (o *WlTouch) ServeShape(id wire.Int, major wire.Fixed, minor wire.Fixed) error {
	if o.shapeHandler == nil {
		return nil
	}
	return o.shapeHandler(id, major, minor)
}

// HandleShape registers a handler for a Shape event.
func (o *WlTouch) HandleShape(h WlTouchShapeHandler) {
	o.shapeHandler = h
}

// ServeOrientation is for update orientation of touch point.
//
// Sent when a touchpoint has changed its orientation.
//
// This event does not occur on its own. It is sent before a
// wl_touch.frame event and carries the new shape information for
// any previously reported, or new touch points of that frame.
//
// Other events describing the touch point such as wl_touch.down,
// wl_touch.motion or wl_touch.shape may be sent within the
// same wl_touch.frame. A client should treat these events as a single
// logical touch point update. The order of wl_touch.shape,
// wl_touch.orientation and wl_touch.motion is not guaranteed.
// A wl_touch.down event is guaranteed to occur before the first
// wl_touch.orientation event for this touch ID but both events may occur
// within the same wl_touch.frame.
//
// The orientation describes the clockwise angle of a touchpoint's major
// axis to the positive surface y-axis and is normalized to the -180 to
// +180 degree range. The granularity of orientation depends on the touch
// device, some devices only support binary rotation values between 0 and
// 90 degrees.
//
// This event is only sent by the compositor if the touch device supports
// orientation reports.
func (o *WlTouch) ServeOrientation(id wire.Int, orientation wire.Fixed) error {
	if o.orientationHandler == nil {
		return nil
	}
	return o.orientationHandler(id, orientation)
}

// HandleOrientation registers a handler for a Orientation event.
func (o *WlTouch) HandleOrientation(h WlTouchOrientationHandler) {
	o.orientationHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlTouch) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.downHandler == nil {
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

		id, err := r.ReadInt()
		if err != nil {
			return err
		}

		x, err := r.ReadFixed()
		if err != nil {
			return err
		}

		y, err := r.ReadFixed()
		if err != nil {
			return err
		}

		return o.downHandler(serial, time, surface, id, x, y)

	case 1:
		if o.upHandler == nil {
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

		id, err := r.ReadInt()
		if err != nil {
			return err
		}

		return o.upHandler(serial, time, id)

	case 2:
		if o.motionHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		time, err := r.ReadUint()
		if err != nil {
			return err
		}

		id, err := r.ReadInt()
		if err != nil {
			return err
		}

		x, err := r.ReadFixed()
		if err != nil {
			return err
		}

		y, err := r.ReadFixed()
		if err != nil {
			return err
		}

		return o.motionHandler(time, id, x, y)

	case 3:
		if o.frameHandler == nil {
			return nil
		}

		return o.frameHandler()

	case 4:
		if o.cancelHandler == nil {
			return nil
		}

		return o.cancelHandler()

	case 5:
		if o.shapeHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		id, err := r.ReadInt()
		if err != nil {
			return err
		}

		major, err := r.ReadFixed()
		if err != nil {
			return err
		}

		minor, err := r.ReadFixed()
		if err != nil {
			return err
		}

		return o.shapeHandler(id, major, minor)

	case 6:
		if o.orientationHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		id, err := r.ReadInt()
		if err != nil {
			return err
		}

		orientation, err := r.ReadFixed()
		if err != nil {
			return err
		}

		return o.orientationHandler(id, orientation)

	default:
		return fmt.Errorf("WlTouch: unhandled message(%v)", msg)
	}
}
