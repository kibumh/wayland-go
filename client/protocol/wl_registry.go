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

// WlRegistryGlobalHandler is a handler for announce global object.
//
// Notify the client of global objects.
//
// The event notifies the client that a global object with
// the given name is now available, and it implements the
// given version of the given interface.
type WlRegistryGlobalHandler func(name wire.Uint, itf wire.String, version wire.Uint) error

// WlRegistryGlobalRemoveHandler is a handler for announce removal of global object.
//
// Notify the client of removed global objects.
//
// This event notifies the client that the global identified
// by name is no longer available.  If the client bound to
// the global using the bind request, the client should now
// destroy that object.
//
// The object remains valid and requests to the object will be
// ignored until the client destroys it, to avoid races between
// the global going away and a client sending a request to it.
type WlRegistryGlobalRemoveHandler func(name wire.Uint) error

/*
 * TYPE
 */
// WlRegistry is global registry object.
//
// The singleton global registry object.  The server has a number of
// global objects that are available to all clients.  These objects
// typically represent an actual object in the server (for example,
// an input device) or they are singleton objects that provide
// extension functionality.
//
// When a client creates a registry object, the registry object
// will emit a global event for each global currently in the
// registry.  Globals come and go as a result of device or
// monitor hotplugs, reconfiguration or other events, and the
// registry will send out global and global_remove events to
// keep the client up to date with the changes.  To mark the end
// of the initial burst of events, the client can use the
// wl_display.sync request immediately after calling
// wl_display.get_registry.
//
// A client can bind to a global object by using the bind
// request.  This creates a client-side handle that lets the object
// emit events to the client and lets the client invoke requests on
// the object.
type WlRegistry struct {
	Base

	globalHandler       WlRegistryGlobalHandler
	globalRemoveHandler WlRegistryGlobalRemoveHandler
}

// NewWlRegistry creates a WlRegistry object.
func NewWlRegistry(c *wire.Conn) *WlRegistry {
	return NewWlRegistryWithID(c, c.NewID())
}

// NewWlRegistryWithID creates a WlRegistry object with a given id.
func NewWlRegistryWithID(c *wire.Conn, id wire.ID) *WlRegistry {
	o := &WlRegistry{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Bind is for bind an object to the display
//
// Binds a new, client-created object to the server using the
// specified name as the identifier.
func (o *WlRegistry) Bind(name wire.Uint, id wire.ID) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(name); err != nil {
		return err
	}

	if err = msg.Write(id); err != nil {
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

// ServeGlobal is for announce global object.
//
// Notify the client of global objects.
//
// The event notifies the client that a global object with
// the given name is now available, and it implements the
// given version of the given interface.
func (o *WlRegistry) ServeGlobal(name wire.Uint, itf wire.String, version wire.Uint) error {
	if o.globalHandler == nil {
		return nil
	}
	return o.globalHandler(name, itf, version)
}

// HandleGlobal registers a handler for a Global event.
func (o *WlRegistry) HandleGlobal(h WlRegistryGlobalHandler) {
	o.globalHandler = h
}

// ServeGlobalRemove is for announce removal of global object.
//
// Notify the client of removed global objects.
//
// This event notifies the client that the global identified
// by name is no longer available.  If the client bound to
// the global using the bind request, the client should now
// destroy that object.
//
// The object remains valid and requests to the object will be
// ignored until the client destroys it, to avoid races between
// the global going away and a client sending a request to it.
func (o *WlRegistry) ServeGlobalRemove(name wire.Uint) error {
	if o.globalRemoveHandler == nil {
		return nil
	}
	return o.globalRemoveHandler(name)
}

// HandleGlobalRemove registers a handler for a GlobalRemove event.
func (o *WlRegistry) HandleGlobalRemove(h WlRegistryGlobalRemoveHandler) {
	o.globalRemoveHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlRegistry) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.globalHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		name, err := r.ReadUint()
		if err != nil {
			return err
		}

		itf, err := r.ReadString()
		if err != nil {
			return err
		}

		version, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.globalHandler(name, itf, version)

	case 1:
		if o.globalRemoveHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		name, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.globalRemoveHandler(name)

	default:
		return fmt.Errorf("WlRegistry: unhandled message(%v)", msg)
	}
}
