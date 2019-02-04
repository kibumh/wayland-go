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

// WlSeatCapability is for seat capability bitmask.
//
// This is a bitmask of capabilities this seat has; if a member is
// set, then it is present on the seat.
type WlSeatCapability int

const (
	WlSeatCapabilityPointer  WlSeatCapability = 1 // the seat has pointer devices
	WlSeatCapabilityKeyboard WlSeatCapability = 2 // the seat has one or more keyboards
	WlSeatCapabilityTouch    WlSeatCapability = 4 // the seat has touch devices
)

/*
 * EVENT HANDLER TYPES
 */

// WlSeatCapabilitiesHandler is a handler for seat capabilities changed.
//
// This is emitted whenever a seat gains or loses the pointer,
// keyboard or touch capabilities.  The argument is a capability
// enum containing the complete set of capabilities this seat has.
//
// When the pointer capability is added, a client may create a
// wl_pointer object using the wl_seat.get_pointer request. This object
// will receive pointer events until the capability is removed in the
// future.
//
// When the pointer capability is removed, a client should destroy the
// wl_pointer objects associated with the seat where the capability was
// removed, using the wl_pointer.release request. No further pointer
// events will be received on these objects.
//
// In some compositors, if a seat regains the pointer capability and a
// client has a previously obtained wl_pointer object of version 4 or
// less, that object may start sending pointer events again. This
// behavior is considered a misinterpretation of the intended behavior
// and must not be relied upon by the client. wl_pointer objects of
// version 5 or later must not send events if created before the most
// recent event notifying the client of an added pointer capability.
//
// The above behavior also applies to wl_keyboard and wl_touch with the
// keyboard and touch capabilities, respectively.
type WlSeatCapabilitiesHandler func(capabilities wire.Uint) error

// WlSeatNameHandler is a handler for unique identifier for this seat.
//
// In a multiseat configuration this can be used by the client to help
// identify which physical devices the seat represents. Based on
// the seat configuration used by the compositor.
type WlSeatNameHandler func(name wire.String) error

/*
 * TYPE
 */
// WlSeat is group of input devices.
//
// A seat is a group of keyboards, pointer and touch devices. This
// object is published as a global during start up, or when such a
// device is hot plugged.  A seat typically has a pointer and
// maintains a keyboard focus and a pointer focus.
type WlSeat struct {
	Base

	capabilitiesHandler WlSeatCapabilitiesHandler
	nameHandler         WlSeatNameHandler
}

// NewWlSeat creates a WlSeat object.
func NewWlSeat(c *wire.Conn) *WlSeat {
	return NewWlSeatWithID(c, c.NewID())
}

// NewWlSeatWithID creates a WlSeat object with a given id.
func NewWlSeatWithID(c *wire.Conn, id wire.ID) *WlSeat {
	o := &WlSeat{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// GetPointer is for return pointer object
//
// The ID provided will be initialized to the wl_pointer interface
// for this seat.
//
// This request only takes effect if the seat has the pointer
// capability, or has had the pointer capability in the past.
// It is a protocol violation to issue this request on a seat that has
// never had the pointer capability.
func (o *WlSeat) GetPointer(id *WlPointer) error {
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

// GetKeyboard is for return keyboard object
//
// The ID provided will be initialized to the wl_keyboard interface
// for this seat.
//
// This request only takes effect if the seat has the keyboard
// capability, or has had the keyboard capability in the past.
// It is a protocol violation to issue this request on a seat that has
// never had the keyboard capability.
func (o *WlSeat) GetKeyboard(id *WlKeyboard) error {
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

// GetTouch is for return touch object
//
// The ID provided will be initialized to the wl_touch interface
// for this seat.
//
// This request only takes effect if the seat has the touch
// capability, or has had the touch capability in the past.
// It is a protocol violation to issue this request on a seat that has
// never had the touch capability.
func (o *WlSeat) GetTouch(id *WlTouch) error {
	msg, err := wire.NewMessage(o.ID(), 2)
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

// Release is for release the seat object
//
// Using this request a client can tell the server that it is not going to
// use the seat object anymore.
func (o *WlSeat) Release() error {
	msg, err := wire.NewMessage(o.ID(), 3)
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

// ServeCapabilities is for seat capabilities changed.
//
// This is emitted whenever a seat gains or loses the pointer,
// keyboard or touch capabilities.  The argument is a capability
// enum containing the complete set of capabilities this seat has.
//
// When the pointer capability is added, a client may create a
// wl_pointer object using the wl_seat.get_pointer request. This object
// will receive pointer events until the capability is removed in the
// future.
//
// When the pointer capability is removed, a client should destroy the
// wl_pointer objects associated with the seat where the capability was
// removed, using the wl_pointer.release request. No further pointer
// events will be received on these objects.
//
// In some compositors, if a seat regains the pointer capability and a
// client has a previously obtained wl_pointer object of version 4 or
// less, that object may start sending pointer events again. This
// behavior is considered a misinterpretation of the intended behavior
// and must not be relied upon by the client. wl_pointer objects of
// version 5 or later must not send events if created before the most
// recent event notifying the client of an added pointer capability.
//
// The above behavior also applies to wl_keyboard and wl_touch with the
// keyboard and touch capabilities, respectively.
func (o *WlSeat) ServeCapabilities(capabilities wire.Uint) error {
	if o.capabilitiesHandler == nil {
		return nil
	}
	return o.capabilitiesHandler(capabilities)
}

// HandleCapabilities registers a handler for a Capabilities event.
func (o *WlSeat) HandleCapabilities(h WlSeatCapabilitiesHandler) {
	o.capabilitiesHandler = h
}

// ServeName is for unique identifier for this seat.
//
// In a multiseat configuration this can be used by the client to help
// identify which physical devices the seat represents. Based on
// the seat configuration used by the compositor.
func (o *WlSeat) ServeName(name wire.String) error {
	if o.nameHandler == nil {
		return nil
	}
	return o.nameHandler(name)
}

// HandleName registers a handler for a Name event.
func (o *WlSeat) HandleName(h WlSeatNameHandler) {
	o.nameHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlSeat) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.capabilitiesHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		capabilities, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.capabilitiesHandler(capabilities)

	case 1:
		if o.nameHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		name, err := r.ReadString()
		if err != nil {
			return err
		}

		return o.nameHandler(name)

	default:
		return fmt.Errorf("WlSeat: unhandled message(%v)", msg)
	}
}
