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

// WlKeyboardKeymapFormat is for keyboard mapping format.
//
// This specifies the format of the keymap provided to the
// client with the wl_keyboard.keymap event.
type WlKeyboardKeymapFormat int

const (
	WlKeyboardKeymapFormatNoKeymap WlKeyboardKeymapFormat = 0 // no keymap; client must understand how to interpret the raw keycode
	WlKeyboardKeymapFormatXkbV1    WlKeyboardKeymapFormat = 1 // libxkbcommon compatible; to determine the xkb keycode, clients must add 8 to the key event keycode
)

// WlKeyboardKeyState is for physical key state.
//
// Describes the physical state of a key that produced the key event.
type WlKeyboardKeyState int

const (
	WlKeyboardKeyStateReleased WlKeyboardKeyState = 0 // key is not pressed
	WlKeyboardKeyStatePressed  WlKeyboardKeyState = 1 // key is pressed
)

/*
 * EVENT HANDLER TYPES
 */

// WlKeyboardKeymapHandler is a handler for keyboard mapping.
//
// This event provides a file descriptor to the client which can be
// memory-mapped to provide a keyboard mapping description.
//
// From version 7 onwards, the fd must be mapped with MAP_PRIVATE by
// the recipient, as MAP_SHARED may fail.
type WlKeyboardKeymapHandler func(format wire.Uint, fd wire.FD, size wire.Uint) error

// WlKeyboardEnterHandler is a handler for enter event.
//
// Notification that this seat's keyboard focus is on a certain
// surface.
type WlKeyboardEnterHandler func(serial wire.Uint, surface *WlSurface, keys wire.Array) error

// WlKeyboardLeaveHandler is a handler for leave event.
//
// Notification that this seat's keyboard focus is no longer on
// a certain surface.
//
// The leave notification is sent before the enter notification
// for the new focus.
type WlKeyboardLeaveHandler func(serial wire.Uint, surface *WlSurface) error

// WlKeyboardKeyHandler is a handler for key event.
//
// A key was pressed or released.
// The time argument is a timestamp with millisecond
// granularity, with an undefined base.
type WlKeyboardKeyHandler func(serial wire.Uint, time wire.Uint, key wire.Uint, state wire.Uint) error

// WlKeyboardModifiersHandler is a handler for modifier and group state.
//
// Notifies clients that the modifier and/or group state has
// changed, and it should update its local state.
type WlKeyboardModifiersHandler func(serial wire.Uint, modsDepressed wire.Uint, modsLatched wire.Uint, modsLocked wire.Uint, group wire.Uint) error

// WlKeyboardRepeatInfoHandler is a handler for repeat rate and delay.
//
// Informs the client about the keyboard's repeat rate and delay.
//
// This event is sent as soon as the wl_keyboard object has been created,
// and is guaranteed to be received by the client before any key press
// event.
//
// Negative values for either rate or delay are illegal. A rate of zero
// will disable any repeating (regardless of the value of delay).
//
// This event can be sent later on as well with a new value if necessary,
// so clients should continue listening for the event past the creation
// of wl_keyboard.
type WlKeyboardRepeatInfoHandler func(rate wire.Int, delay wire.Int) error

/*
 * TYPE
 */
// WlKeyboard is keyboard input device.
//
// The wl_keyboard interface represents one or more keyboards
// associated with a seat.
type WlKeyboard struct {
	Base

	keymapHandler     WlKeyboardKeymapHandler
	enterHandler      WlKeyboardEnterHandler
	leaveHandler      WlKeyboardLeaveHandler
	keyHandler        WlKeyboardKeyHandler
	modifiersHandler  WlKeyboardModifiersHandler
	repeatInfoHandler WlKeyboardRepeatInfoHandler
}

// NewWlKeyboard creates a WlKeyboard object.
func NewWlKeyboard(c *wire.Conn) *WlKeyboard {
	return NewWlKeyboardWithID(c, c.NewID())
}

// NewWlKeyboardWithID creates a WlKeyboard object with a given id.
func NewWlKeyboardWithID(c *wire.Conn, id wire.ID) *WlKeyboard {
	o := &WlKeyboard{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Release is for release the keyboard object
//
//
func (o *WlKeyboard) Release() error {
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

// ServeKeymap is for keyboard mapping.
//
// This event provides a file descriptor to the client which can be
// memory-mapped to provide a keyboard mapping description.
//
// From version 7 onwards, the fd must be mapped with MAP_PRIVATE by
// the recipient, as MAP_SHARED may fail.
func (o *WlKeyboard) ServeKeymap(format wire.Uint, fd wire.FD, size wire.Uint) error {
	if o.keymapHandler == nil {
		return nil
	}
	return o.keymapHandler(format, fd, size)
}

// HandleKeymap registers a handler for a Keymap event.
func (o *WlKeyboard) HandleKeymap(h WlKeyboardKeymapHandler) {
	o.keymapHandler = h
}

// ServeEnter is for enter event.
//
// Notification that this seat's keyboard focus is on a certain
// surface.
func (o *WlKeyboard) ServeEnter(serial wire.Uint, surface *WlSurface, keys wire.Array) error {
	if o.enterHandler == nil {
		return nil
	}
	return o.enterHandler(serial, surface, keys)
}

// HandleEnter registers a handler for a Enter event.
func (o *WlKeyboard) HandleEnter(h WlKeyboardEnterHandler) {
	o.enterHandler = h
}

// ServeLeave is for leave event.
//
// Notification that this seat's keyboard focus is no longer on
// a certain surface.
//
// The leave notification is sent before the enter notification
// for the new focus.
func (o *WlKeyboard) ServeLeave(serial wire.Uint, surface *WlSurface) error {
	if o.leaveHandler == nil {
		return nil
	}
	return o.leaveHandler(serial, surface)
}

// HandleLeave registers a handler for a Leave event.
func (o *WlKeyboard) HandleLeave(h WlKeyboardLeaveHandler) {
	o.leaveHandler = h
}

// ServeKey is for key event.
//
// A key was pressed or released.
// The time argument is a timestamp with millisecond
// granularity, with an undefined base.
func (o *WlKeyboard) ServeKey(serial wire.Uint, time wire.Uint, key wire.Uint, state wire.Uint) error {
	if o.keyHandler == nil {
		return nil
	}
	return o.keyHandler(serial, time, key, state)
}

// HandleKey registers a handler for a Key event.
func (o *WlKeyboard) HandleKey(h WlKeyboardKeyHandler) {
	o.keyHandler = h
}

// ServeModifiers is for modifier and group state.
//
// Notifies clients that the modifier and/or group state has
// changed, and it should update its local state.
func (o *WlKeyboard) ServeModifiers(serial wire.Uint, modsDepressed wire.Uint, modsLatched wire.Uint, modsLocked wire.Uint, group wire.Uint) error {
	if o.modifiersHandler == nil {
		return nil
	}
	return o.modifiersHandler(serial, modsDepressed, modsLatched, modsLocked, group)
}

// HandleModifiers registers a handler for a Modifiers event.
func (o *WlKeyboard) HandleModifiers(h WlKeyboardModifiersHandler) {
	o.modifiersHandler = h
}

// ServeRepeatInfo is for repeat rate and delay.
//
// Informs the client about the keyboard's repeat rate and delay.
//
// This event is sent as soon as the wl_keyboard object has been created,
// and is guaranteed to be received by the client before any key press
// event.
//
// Negative values for either rate or delay are illegal. A rate of zero
// will disable any repeating (regardless of the value of delay).
//
// This event can be sent later on as well with a new value if necessary,
// so clients should continue listening for the event past the creation
// of wl_keyboard.
func (o *WlKeyboard) ServeRepeatInfo(rate wire.Int, delay wire.Int) error {
	if o.repeatInfoHandler == nil {
		return nil
	}
	return o.repeatInfoHandler(rate, delay)
}

// HandleRepeatInfo registers a handler for a RepeatInfo event.
func (o *WlKeyboard) HandleRepeatInfo(h WlKeyboardRepeatInfoHandler) {
	o.repeatInfoHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlKeyboard) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.keymapHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		format, err := r.ReadUint()
		if err != nil {
			return err
		}

		fd, err := r.ReadFD()
		if err != nil {
			return err
		}

		size, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.keymapHandler(format, fd, size)

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

		keys, err := r.ReadArray()
		if err != nil {
			return err
		}

		return o.enterHandler(serial, surface, keys)

	case 2:
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

	case 3:
		if o.keyHandler == nil {
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

		key, err := r.ReadUint()
		if err != nil {
			return err
		}

		state, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.keyHandler(serial, time, key, state)

	case 4:
		if o.modifiersHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		serial, err := r.ReadUint()
		if err != nil {
			return err
		}

		modsDepressed, err := r.ReadUint()
		if err != nil {
			return err
		}

		modsLatched, err := r.ReadUint()
		if err != nil {
			return err
		}

		modsLocked, err := r.ReadUint()
		if err != nil {
			return err
		}

		group, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.modifiersHandler(serial, modsDepressed, modsLatched, modsLocked, group)

	case 5:
		if o.repeatInfoHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		rate, err := r.ReadInt()
		if err != nil {
			return err
		}

		delay, err := r.ReadInt()
		if err != nil {
			return err
		}

		return o.repeatInfoHandler(rate, delay)

	default:
		return fmt.Errorf("WlKeyboard: unhandled message(%v)", msg)
	}
}
