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

// WlDataDeviceManagerDndAction is for drag and drop actions.
//
// This is a bitmask of the available/preferred actions in a
// drag-and-drop operation.
//
// In the compositor, the selected action is a result of matching the
// actions offered by the source and destination sides.  "action" events
// with a "none" action will be sent to both source and destination if
// there is no match. All further checks will effectively happen on
// (source actions ∩ destination actions).
//
// In addition, compositors may also pick different actions in
// reaction to key modifiers being pressed. One common design that
// is used in major toolkits (and the behavior recommended for
// compositors) is:
//
// - If no modifiers are pressed, the first match (in bit order)
// will be used.
// - Pressing Shift selects "move", if enabled in the mask.
// - Pressing Control selects "copy", if enabled in the mask.
//
// Behavior beyond that is considered implementation-dependent.
// Compositors may for example bind other modifiers (like Alt/Meta)
// or drags initiated with other buttons than BTN_LEFT to specific
// actions (e.g. "ask").
type WlDataDeviceManagerDndAction int

const (
	WlDataDeviceManagerDndActionNone WlDataDeviceManagerDndAction = 0 // no action
	WlDataDeviceManagerDndActionCopy WlDataDeviceManagerDndAction = 1 // copy action
	WlDataDeviceManagerDndActionMove WlDataDeviceManagerDndAction = 2 // move action
	WlDataDeviceManagerDndActionAsk  WlDataDeviceManagerDndAction = 4 // ask action
)

/*
 * EVENT HANDLER TYPES
 */

/*
 * TYPE
 */
// WlDataDeviceManager is data transfer interface.
//
// The wl_data_device_manager is a singleton global object that
// provides access to inter-client data transfer mechanisms such as
// copy-and-paste and drag-and-drop.  These mechanisms are tied to
// a wl_seat and this interface lets a client get a wl_data_device
// corresponding to a wl_seat.
//
// Depending on the version bound, the objects created from the bound
// wl_data_device_manager object will have different requirements for
// functioning properly. See wl_data_source.set_actions,
// wl_data_offer.accept and wl_data_offer.finish for details.
type WlDataDeviceManager struct {
	Base
}

// NewWlDataDeviceManager creates a WlDataDeviceManager object.
func NewWlDataDeviceManager(c *wire.Conn) *WlDataDeviceManager {
	return NewWlDataDeviceManagerWithID(c, c.NewID())
}

// NewWlDataDeviceManagerWithID creates a WlDataDeviceManager object with a given id.
func NewWlDataDeviceManagerWithID(c *wire.Conn, id wire.ID) *WlDataDeviceManager {
	o := &WlDataDeviceManager{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// CreateDataSource is for create a new data source
//
// Create a new data source.
func (o *WlDataDeviceManager) CreateDataSource(id *WlDataSource) error {
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

// GetDataDevice is for create a new data device
//
// Create a new data device for a given seat.
func (o *WlDataDeviceManager) GetDataDevice(id *WlDataDevice, seat *WlSeat) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(id.ID()); err != nil {
		return err
	}

	if err = msg.Write(seat.ID()); err != nil {
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

// ServeMessage is a multiplexer for a message.
func (o *WlDataDeviceManager) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	default:
		return fmt.Errorf("WlDataDeviceManager: unhandled message(%v)", msg)
	}
}
