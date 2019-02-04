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

// WlDisplayError is for global error values.
//
// These errors are global and can be emitted in response to any
// server request.
type WlDisplayError int

const (
	WlDisplayErrorInvalidObject WlDisplayError = 0 // server couldn't find object
	WlDisplayErrorInvalidMethod WlDisplayError = 1 // method doesn't exist on the specified interface
	WlDisplayErrorNoMemory      WlDisplayError = 2 // server is out of memory
)

/*
 * EVENT HANDLER TYPES
 */

// WlDisplayErrorHandler is a handler for fatal error event.
//
// The error event is sent out when a fatal (non-recoverable)
// error has occurred.  The object_id argument is the object
// where the error occurred, most often in response to a request
// to that object.  The code identifies the error and is defined
// by the object interface.  As such, each interface defines its
// own set of error codes.  The message is a brief description
// of the error, for (debugging) convenience.
type WlDisplayErrorHandler func(objectId wire.Object, code wire.Uint, message wire.String) error

// WlDisplayDeleteIdHandler is a handler for acknowledge object ID deletion.
//
// This event is used internally by the object ID management
// logic.  When a client deletes an object, the server will send
// this event to acknowledge that it has seen the delete request.
// When the client receives this event, it will know that it can
// safely reuse the object ID.
type WlDisplayDeleteIdHandler func(id wire.Uint) error

/*
 * TYPE
 */
// WlDisplay is core global object.
//
// The core global object.  This is a special singleton object.  It
// is used for internal Wayland protocol features.
type WlDisplay struct {
	Base

	errorHandler    WlDisplayErrorHandler
	deleteIdHandler WlDisplayDeleteIdHandler
}

// NewWlDisplay creates a WlDisplay object.
func NewWlDisplay(c *wire.Conn) *WlDisplay {
	return NewWlDisplayWithID(c, c.NewID())
}

// NewWlDisplayWithID creates a WlDisplay object with a given id.
func NewWlDisplayWithID(c *wire.Conn, id wire.ID) *WlDisplay {
	o := &WlDisplay{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Sync is for asynchronous roundtrip
//
// The sync request asks the server to emit the 'done' event
// on the returned wl_callback object.  Since requests are
// handled in-order and events are delivered in-order, this can
// be used as a barrier to ensure all previous requests and the
// resulting events have been handled.
//
// The object returned by this request will be destroyed by the
// compositor after the callback is fired and as such the client must not
// attempt to use it after that point.
//
// The callback_data passed in the callback is the event serial.
func (o *WlDisplay) Sync(callback *WlCallback) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(callback.ID()); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// GetRegistry is for get global registry object
//
// This request creates a registry object that allows the client
// to list and bind the global objects available from the
// compositor.
//
// It should be noted that the server side resources consumed in
// response to a get_registry request can only be released when the
// client disconnects, not when the client side proxy is destroyed.
// Therefore, clients should invoke get_registry as infrequently as
// possible to avoid wasting memory.
func (o *WlDisplay) GetRegistry(registry *WlRegistry) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(registry.ID()); err != nil {
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

// ServeError is for fatal error event.
//
// The error event is sent out when a fatal (non-recoverable)
// error has occurred.  The object_id argument is the object
// where the error occurred, most often in response to a request
// to that object.  The code identifies the error and is defined
// by the object interface.  As such, each interface defines its
// own set of error codes.  The message is a brief description
// of the error, for (debugging) convenience.
func (o *WlDisplay) ServeError(objectId wire.Object, code wire.Uint, message wire.String) error {
	if o.errorHandler == nil {
		return nil
	}
	return o.errorHandler(objectId, code, message)
}

// HandleError registers a handler for a Error event.
func (o *WlDisplay) HandleError(h WlDisplayErrorHandler) {
	o.errorHandler = h
}

// ServeDeleteId is for acknowledge object ID deletion.
//
// This event is used internally by the object ID management
// logic.  When a client deletes an object, the server will send
// this event to acknowledge that it has seen the delete request.
// When the client receives this event, it will know that it can
// safely reuse the object ID.
func (o *WlDisplay) ServeDeleteId(id wire.Uint) error {
	if o.deleteIdHandler == nil {
		return nil
	}
	return o.deleteIdHandler(id)
}

// HandleDeleteId registers a handler for a DeleteId event.
func (o *WlDisplay) HandleDeleteId(h WlDisplayDeleteIdHandler) {
	o.deleteIdHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlDisplay) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.errorHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		objectIdId, err := r.ReadID()
		if err != nil {
			return err
		}

		objectId, ok := o.Base.Conn.GetObject(objectIdId)
		if !ok {
			return fmt.Errorf("cannot find an object: id(%d)", objectIdId)
		}
		code, err := r.ReadUint()
		if err != nil {
			return err
		}

		message, err := r.ReadString()
		if err != nil {
			return err
		}

		return o.errorHandler(objectId, code, message)

	case 1:
		if o.deleteIdHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		id, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.deleteIdHandler(id)

	default:
		return fmt.Errorf("WlDisplay: unhandled message(%v)", msg)
	}
}
