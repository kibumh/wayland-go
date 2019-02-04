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

// WlCallbackDoneHandler is a handler for done event.
//
// Notify the client when the related request is done.
type WlCallbackDoneHandler func(callbackData wire.Uint) error

/*
 * TYPE
 */
// WlCallback is callback object.
//
// Clients can handle the 'done' event to get notified when
// the related request is done.
type WlCallback struct {
	Base

	doneHandler WlCallbackDoneHandler
}

// NewWlCallback creates a WlCallback object.
func NewWlCallback(c *wire.Conn) *WlCallback {
	return NewWlCallbackWithID(c, c.NewID())
}

// NewWlCallbackWithID creates a WlCallback object with a given id.
func NewWlCallbackWithID(c *wire.Conn, id wire.ID) *WlCallback {
	o := &WlCallback{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

/*
 * EVENTS
 */

// ServeDone is for done event.
//
// Notify the client when the related request is done.
func (o *WlCallback) ServeDone(callbackData wire.Uint) error {
	if o.doneHandler == nil {
		return nil
	}
	return o.doneHandler(callbackData)
}

// HandleDone registers a handler for a Done event.
func (o *WlCallback) HandleDone(h WlCallbackDoneHandler) {
	o.doneHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlCallback) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.doneHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		callbackData, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.doneHandler(callbackData)

	default:
		return fmt.Errorf("WlCallback: unhandled message(%v)", msg)
	}
}
