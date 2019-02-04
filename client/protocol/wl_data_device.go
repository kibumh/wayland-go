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

// WlDataDeviceError is for .
//
//
type WlDataDeviceError int

const (
	WlDataDeviceErrorRole WlDataDeviceError = 0 // given wl_surface has another role
)

/*
 * EVENT HANDLER TYPES
 */

// WlDataDeviceDataOfferHandler is a handler for introduce a new wl_data_offer.
//
// The data_offer event introduces a new wl_data_offer object,
// which will subsequently be used in either the
// data_device.enter event (for drag-and-drop) or the
// data_device.selection event (for selections).  Immediately
// following the data_device_data_offer event, the new data_offer
// object will send out data_offer.offer events to describe the
// mime types it offers.
type WlDataDeviceDataOfferHandler func(id *WlDataOffer) error

// WlDataDeviceEnterHandler is a handler for initiate drag-and-drop session.
//
// This event is sent when an active drag-and-drop pointer enters
// a surface owned by the client.  The position of the pointer at
// enter time is provided by the x and y arguments, in surface-local
// coordinates.
type WlDataDeviceEnterHandler func(serial wire.Uint, surface *WlSurface, x wire.Fixed, y wire.Fixed, id *WlDataOffer) error

// WlDataDeviceLeaveHandler is a handler for end drag-and-drop session.
//
// This event is sent when the drag-and-drop pointer leaves the
// surface and the session ends.  The client must destroy the
// wl_data_offer introduced at enter time at this point.
type WlDataDeviceLeaveHandler func() error

// WlDataDeviceMotionHandler is a handler for drag-and-drop session motion.
//
// This event is sent when the drag-and-drop pointer moves within
// the currently focused surface. The new position of the pointer
// is provided by the x and y arguments, in surface-local
// coordinates.
type WlDataDeviceMotionHandler func(time wire.Uint, x wire.Fixed, y wire.Fixed) error

// WlDataDeviceDropHandler is a handler for end drag-and-drop session successfully.
//
// The event is sent when a drag-and-drop operation is ended
// because the implicit grab is removed.
//
// The drag-and-drop destination is expected to honor the last action
// received through wl_data_offer.action, if the resulting action is
// "copy" or "move", the destination can still perform
// wl_data_offer.receive requests, and is expected to end all
// transfers with a wl_data_offer.finish request.
//
// If the resulting action is "ask", the action will not be considered
// final. The drag-and-drop destination is expected to perform one last
// wl_data_offer.set_actions request, or wl_data_offer.destroy in order
// to cancel the operation.
type WlDataDeviceDropHandler func() error

// WlDataDeviceSelectionHandler is a handler for advertise new selection.
//
// The selection event is sent out to notify the client of a new
// wl_data_offer for the selection for this device.  The
// data_device.data_offer and the data_offer.offer events are
// sent out immediately before this event to introduce the data
// offer object.  The selection event is sent to a client
// immediately before receiving keyboard focus and when a new
// selection is set while the client has keyboard focus.  The
// data_offer is valid until a new data_offer or NULL is received
// or until the client loses keyboard focus.  The client must
// destroy the previous selection data_offer, if any, upon receiving
// this event.
type WlDataDeviceSelectionHandler func(id *WlDataOffer) error

/*
 * TYPE
 */
// WlDataDevice is data transfer device.
//
// There is one wl_data_device per seat which can be obtained
// from the global wl_data_device_manager singleton.
//
// A wl_data_device provides access to inter-client data transfer
// mechanisms such as copy-and-paste and drag-and-drop.
type WlDataDevice struct {
	Base

	dataOfferHandler WlDataDeviceDataOfferHandler
	enterHandler     WlDataDeviceEnterHandler
	leaveHandler     WlDataDeviceLeaveHandler
	motionHandler    WlDataDeviceMotionHandler
	dropHandler      WlDataDeviceDropHandler
	selectionHandler WlDataDeviceSelectionHandler
}

// NewWlDataDevice creates a WlDataDevice object.
func NewWlDataDevice(c *wire.Conn) *WlDataDevice {
	return NewWlDataDeviceWithID(c, c.NewID())
}

// NewWlDataDeviceWithID creates a WlDataDevice object with a given id.
func NewWlDataDeviceWithID(c *wire.Conn, id wire.ID) *WlDataDevice {
	o := &WlDataDevice{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// StartDrag is for start drag-and-drop operation
//
// This request asks the compositor to start a drag-and-drop
// operation on behalf of the client.
//
// The source argument is the data source that provides the data
// for the eventual data transfer. If source is NULL, enter, leave
// and motion events are sent only to the client that initiated the
// drag and the client is expected to handle the data passing
// internally.
//
// The origin surface is the surface where the drag originates and
// the client must have an active implicit grab that matches the
// serial.
//
// The icon surface is an optional (can be NULL) surface that
// provides an icon to be moved around with the cursor.  Initially,
// the top-left corner of the icon surface is placed at the cursor
// hotspot, but subsequent wl_surface.attach request can move the
// relative position. Attach requests must be confirmed with
// wl_surface.commit as usual. The icon surface is given the role of
// a drag-and-drop icon. If the icon surface already has another role,
// it raises a protocol error.
//
// The current and pending input regions of the icon wl_surface are
// cleared, and wl_surface.set_input_region is ignored until the
// wl_surface is no longer used as the icon surface. When the use
// as an icon ends, the current and pending input regions become
// undefined, and the wl_surface is unmapped.
func (o *WlDataDevice) StartDrag(source *WlDataSource, origin *WlSurface, icon *WlSurface, serial wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(source.ID()); err != nil {
		return err
	}

	if err = msg.Write(origin.ID()); err != nil {
		return err
	}

	if err = msg.Write(icon.ID()); err != nil {
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

// SetSelection is for copy data to the selection
//
// This request asks the compositor to set the selection
// to the data from the source on behalf of the client.
//
// To unset the selection, set the source to NULL.
func (o *WlDataDevice) SetSelection(source *WlDataSource, serial wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(source.ID()); err != nil {
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

// Release is for destroy data device
//
// This request destroys the data device.
func (o *WlDataDevice) Release() error {
	msg, err := wire.NewMessage(o.ID(), 2)
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

// ServeDataOffer is for introduce a new wl_data_offer.
//
// The data_offer event introduces a new wl_data_offer object,
// which will subsequently be used in either the
// data_device.enter event (for drag-and-drop) or the
// data_device.selection event (for selections).  Immediately
// following the data_device_data_offer event, the new data_offer
// object will send out data_offer.offer events to describe the
// mime types it offers.
func (o *WlDataDevice) ServeDataOffer(id *WlDataOffer) error {
	if o.dataOfferHandler == nil {
		return nil
	}
	return o.dataOfferHandler(id)
}

// HandleDataOffer registers a handler for a DataOffer event.
func (o *WlDataDevice) HandleDataOffer(h WlDataDeviceDataOfferHandler) {
	o.dataOfferHandler = h
}

// ServeEnter is for initiate drag-and-drop session.
//
// This event is sent when an active drag-and-drop pointer enters
// a surface owned by the client.  The position of the pointer at
// enter time is provided by the x and y arguments, in surface-local
// coordinates.
func (o *WlDataDevice) ServeEnter(serial wire.Uint, surface *WlSurface, x wire.Fixed, y wire.Fixed, id *WlDataOffer) error {
	if o.enterHandler == nil {
		return nil
	}
	return o.enterHandler(serial, surface, x, y, id)
}

// HandleEnter registers a handler for a Enter event.
func (o *WlDataDevice) HandleEnter(h WlDataDeviceEnterHandler) {
	o.enterHandler = h
}

// ServeLeave is for end drag-and-drop session.
//
// This event is sent when the drag-and-drop pointer leaves the
// surface and the session ends.  The client must destroy the
// wl_data_offer introduced at enter time at this point.
func (o *WlDataDevice) ServeLeave() error {
	if o.leaveHandler == nil {
		return nil
	}
	return o.leaveHandler()
}

// HandleLeave registers a handler for a Leave event.
func (o *WlDataDevice) HandleLeave(h WlDataDeviceLeaveHandler) {
	o.leaveHandler = h
}

// ServeMotion is for drag-and-drop session motion.
//
// This event is sent when the drag-and-drop pointer moves within
// the currently focused surface. The new position of the pointer
// is provided by the x and y arguments, in surface-local
// coordinates.
func (o *WlDataDevice) ServeMotion(time wire.Uint, x wire.Fixed, y wire.Fixed) error {
	if o.motionHandler == nil {
		return nil
	}
	return o.motionHandler(time, x, y)
}

// HandleMotion registers a handler for a Motion event.
func (o *WlDataDevice) HandleMotion(h WlDataDeviceMotionHandler) {
	o.motionHandler = h
}

// ServeDrop is for end drag-and-drop session successfully.
//
// The event is sent when a drag-and-drop operation is ended
// because the implicit grab is removed.
//
// The drag-and-drop destination is expected to honor the last action
// received through wl_data_offer.action, if the resulting action is
// "copy" or "move", the destination can still perform
// wl_data_offer.receive requests, and is expected to end all
// transfers with a wl_data_offer.finish request.
//
// If the resulting action is "ask", the action will not be considered
// final. The drag-and-drop destination is expected to perform one last
// wl_data_offer.set_actions request, or wl_data_offer.destroy in order
// to cancel the operation.
func (o *WlDataDevice) ServeDrop() error {
	if o.dropHandler == nil {
		return nil
	}
	return o.dropHandler()
}

// HandleDrop registers a handler for a Drop event.
func (o *WlDataDevice) HandleDrop(h WlDataDeviceDropHandler) {
	o.dropHandler = h
}

// ServeSelection is for advertise new selection.
//
// The selection event is sent out to notify the client of a new
// wl_data_offer for the selection for this device.  The
// data_device.data_offer and the data_offer.offer events are
// sent out immediately before this event to introduce the data
// offer object.  The selection event is sent to a client
// immediately before receiving keyboard focus and when a new
// selection is set while the client has keyboard focus.  The
// data_offer is valid until a new data_offer or NULL is received
// or until the client loses keyboard focus.  The client must
// destroy the previous selection data_offer, if any, upon receiving
// this event.
func (o *WlDataDevice) ServeSelection(id *WlDataOffer) error {
	if o.selectionHandler == nil {
		return nil
	}
	return o.selectionHandler(id)
}

// HandleSelection registers a handler for a Selection event.
func (o *WlDataDevice) HandleSelection(h WlDataDeviceSelectionHandler) {
	o.selectionHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlDataDevice) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.dataOfferHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		idId, err := r.ReadID()
		if err != nil {
			return err
		}

		id := NewWlDataOfferWithID(o.Base.Conn, idId)

		return o.dataOfferHandler(id)

	case 1:
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

		x, err := r.ReadFixed()
		if err != nil {
			return err
		}

		y, err := r.ReadFixed()
		if err != nil {
			return err
		}

		idId, err := r.ReadID()
		if err != nil {
			return err
		}

		idObj, ok := o.Base.Conn.GetObject(idId)
		if !ok {
			return fmt.Errorf("cannot find an object: id(%d)", idId)
		}
		id, ok := idObj.(*WlDataOffer)
		if !ok {
			return fmt.Errorf("failed to type assertion: id(%d), type(WlDataOffer)", idId)
		}

		return o.enterHandler(serial, surface, x, y, id)

	case 2:
		if o.leaveHandler == nil {
			return nil
		}

		return o.leaveHandler()

	case 3:
		if o.motionHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		time, err := r.ReadUint()
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

		return o.motionHandler(time, x, y)

	case 4:
		if o.dropHandler == nil {
			return nil
		}

		return o.dropHandler()

	case 5:
		if o.selectionHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		idId, err := r.ReadID()
		if err != nil {
			return err
		}

		idObj, ok := o.Base.Conn.GetObject(idId)
		if !ok {
			return fmt.Errorf("cannot find an object: id(%d)", idId)
		}
		id, ok := idObj.(*WlDataOffer)
		if !ok {
			return fmt.Errorf("failed to type assertion: id(%d), type(WlDataOffer)", idId)
		}

		return o.selectionHandler(id)

	default:
		return fmt.Errorf("WlDataDevice: unhandled message(%v)", msg)
	}
}
