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

// WlDataSourceError is for .
//
//
type WlDataSourceError int

const (
	WlDataSourceErrorInvalidActionMask WlDataSourceError = 0 // action mask contains invalid values
	WlDataSourceErrorInvalidSource     WlDataSourceError = 1 // source doesn't accept this request
)

/*
 * EVENT HANDLER TYPES
 */

// WlDataSourceTargetHandler is a handler for a target accepts an offered mime type.
//
// Sent when a target accepts pointer_focus or motion events.  If
// a target does not accept any of the offered types, type is NULL.
//
// Used for feedback during drag-and-drop.
type WlDataSourceTargetHandler func(mimeType wire.String) error

// WlDataSourceSendHandler is a handler for send the data.
//
// Request for data from the client.  Send the data as the
// specified mime type over the passed file descriptor, then
// close it.
type WlDataSourceSendHandler func(mimeType wire.String, fd wire.FD) error

// WlDataSourceCancelledHandler is a handler for selection was cancelled.
//
// This data source is no longer valid. There are several reasons why
// this could happen:
//
// - The data source has been replaced by another data source.
// - The drag-and-drop operation was performed, but the drop destination
// did not accept any of the mime types offered through
// wl_data_source.target.
// - The drag-and-drop operation was performed, but the drop destination
// did not select any of the actions present in the mask offered through
// wl_data_source.action.
// - The drag-and-drop operation was performed but didn't happen over a
// surface.
// - The compositor cancelled the drag-and-drop operation (e.g. compositor
// dependent timeouts to avoid stale drag-and-drop transfers).
//
// The client should clean up and destroy this data source.
//
// For objects of version 2 or older, wl_data_source.cancelled will
// only be emitted if the data source was replaced by another data
// source.
type WlDataSourceCancelledHandler func() error

// WlDataSourceDndDropPerformedHandler is a handler for the drag-and-drop operation physically finished.
//
// The user performed the drop action. This event does not indicate
// acceptance, wl_data_source.cancelled may still be emitted afterwards
// if the drop destination does not accept any mime type.
//
// However, this event might however not be received if the compositor
// cancelled the drag-and-drop operation before this event could happen.
//
// Note that the data_source may still be used in the future and should
// not be destroyed here.
type WlDataSourceDndDropPerformedHandler func() error

// WlDataSourceDndFinishedHandler is a handler for the drag-and-drop operation concluded.
//
// The drop destination finished interoperating with this data
// source, so the client is now free to destroy this data source and
// free all associated data.
//
// If the action used to perform the operation was "move", the
// source can now delete the transferred data.
type WlDataSourceDndFinishedHandler func() error

// WlDataSourceActionHandler is a handler for notify the selected action.
//
// This event indicates the action selected by the compositor after
// matching the source/destination side actions. Only one action (or
// none) will be offered here.
//
// This event can be emitted multiple times during the drag-and-drop
// operation, mainly in response to destination side changes through
// wl_data_offer.set_actions, and as the data device enters/leaves
// surfaces.
//
// It is only possible to receive this event after
// wl_data_source.dnd_drop_performed if the drag-and-drop operation
// ended in an "ask" action, in which case the final wl_data_source.action
// event will happen immediately before wl_data_source.dnd_finished.
//
// Compositors may also change the selected action on the fly, mainly
// in response to keyboard modifier changes during the drag-and-drop
// operation.
//
// The most recent action received is always the valid one. The chosen
// action may change alongside negotiation (e.g. an "ask" action can turn
// into a "move" operation), so the effects of the final action must
// always be applied in wl_data_offer.dnd_finished.
//
// Clients can trigger cursor surface changes from this point, so
// they reflect the current action.
type WlDataSourceActionHandler func(dndAction wire.Uint) error

/*
 * TYPE
 */
// WlDataSource is offer to transfer data.
//
// The wl_data_source object is the source side of a wl_data_offer.
// It is created by the source client in a data transfer and
// provides a way to describe the offered data and a way to respond
// to requests to transfer the data.
type WlDataSource struct {
	Base

	targetHandler           WlDataSourceTargetHandler
	sendHandler             WlDataSourceSendHandler
	cancelledHandler        WlDataSourceCancelledHandler
	dndDropPerformedHandler WlDataSourceDndDropPerformedHandler
	dndFinishedHandler      WlDataSourceDndFinishedHandler
	actionHandler           WlDataSourceActionHandler
}

// NewWlDataSource creates a WlDataSource object.
func NewWlDataSource(c *wire.Conn) *WlDataSource {
	return NewWlDataSourceWithID(c, c.NewID())
}

// NewWlDataSourceWithID creates a WlDataSource object with a given id.
func NewWlDataSourceWithID(c *wire.Conn, id wire.ID) *WlDataSource {
	o := &WlDataSource{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Offer is for add an offered mime type
//
// This request adds a mime type to the set of mime types
// advertised to targets.  Can be called several times to offer
// multiple types.
func (o *WlDataSource) Offer(mimeType wire.String) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(mimeType); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Destroy is for destroy the data source
//
// Destroy the data source.
func (o *WlDataSource) Destroy() error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetActions is for set the available drag-and-drop actions
//
// Sets the actions that the source side client supports for this
// operation. This request may trigger wl_data_source.action and
// wl_data_offer.action events if the compositor needs to change the
// selected action.
//
// The dnd_actions argument must contain only values expressed in the
// wl_data_device_manager.dnd_actions enum, otherwise it will result
// in a protocol error.
//
// This request must be made once only, and can only be made on sources
// used in drag-and-drop, so it must be performed before
// wl_data_device.start_drag. Attempting to use the source other than
// for drag-and-drop will raise a protocol error.
func (o *WlDataSource) SetActions(dndActions wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 2)
	if err != nil {
		return err
	}

	if err = msg.Write(dndActions); err != nil {
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

// ServeTarget is for a target accepts an offered mime type.
//
// Sent when a target accepts pointer_focus or motion events.  If
// a target does not accept any of the offered types, type is NULL.
//
// Used for feedback during drag-and-drop.
func (o *WlDataSource) ServeTarget(mimeType wire.String) error {
	if o.targetHandler == nil {
		return nil
	}
	return o.targetHandler(mimeType)
}

// HandleTarget registers a handler for a Target event.
func (o *WlDataSource) HandleTarget(h WlDataSourceTargetHandler) {
	o.targetHandler = h
}

// ServeSend is for send the data.
//
// Request for data from the client.  Send the data as the
// specified mime type over the passed file descriptor, then
// close it.
func (o *WlDataSource) ServeSend(mimeType wire.String, fd wire.FD) error {
	if o.sendHandler == nil {
		return nil
	}
	return o.sendHandler(mimeType, fd)
}

// HandleSend registers a handler for a Send event.
func (o *WlDataSource) HandleSend(h WlDataSourceSendHandler) {
	o.sendHandler = h
}

// ServeCancelled is for selection was cancelled.
//
// This data source is no longer valid. There are several reasons why
// this could happen:
//
// - The data source has been replaced by another data source.
// - The drag-and-drop operation was performed, but the drop destination
// did not accept any of the mime types offered through
// wl_data_source.target.
// - The drag-and-drop operation was performed, but the drop destination
// did not select any of the actions present in the mask offered through
// wl_data_source.action.
// - The drag-and-drop operation was performed but didn't happen over a
// surface.
// - The compositor cancelled the drag-and-drop operation (e.g. compositor
// dependent timeouts to avoid stale drag-and-drop transfers).
//
// The client should clean up and destroy this data source.
//
// For objects of version 2 or older, wl_data_source.cancelled will
// only be emitted if the data source was replaced by another data
// source.
func (o *WlDataSource) ServeCancelled() error {
	if o.cancelledHandler == nil {
		return nil
	}
	return o.cancelledHandler()
}

// HandleCancelled registers a handler for a Cancelled event.
func (o *WlDataSource) HandleCancelled(h WlDataSourceCancelledHandler) {
	o.cancelledHandler = h
}

// ServeDndDropPerformed is for the drag-and-drop operation physically finished.
//
// The user performed the drop action. This event does not indicate
// acceptance, wl_data_source.cancelled may still be emitted afterwards
// if the drop destination does not accept any mime type.
//
// However, this event might however not be received if the compositor
// cancelled the drag-and-drop operation before this event could happen.
//
// Note that the data_source may still be used in the future and should
// not be destroyed here.
func (o *WlDataSource) ServeDndDropPerformed() error {
	if o.dndDropPerformedHandler == nil {
		return nil
	}
	return o.dndDropPerformedHandler()
}

// HandleDndDropPerformed registers a handler for a DndDropPerformed event.
func (o *WlDataSource) HandleDndDropPerformed(h WlDataSourceDndDropPerformedHandler) {
	o.dndDropPerformedHandler = h
}

// ServeDndFinished is for the drag-and-drop operation concluded.
//
// The drop destination finished interoperating with this data
// source, so the client is now free to destroy this data source and
// free all associated data.
//
// If the action used to perform the operation was "move", the
// source can now delete the transferred data.
func (o *WlDataSource) ServeDndFinished() error {
	if o.dndFinishedHandler == nil {
		return nil
	}
	return o.dndFinishedHandler()
}

// HandleDndFinished registers a handler for a DndFinished event.
func (o *WlDataSource) HandleDndFinished(h WlDataSourceDndFinishedHandler) {
	o.dndFinishedHandler = h
}

// ServeAction is for notify the selected action.
//
// This event indicates the action selected by the compositor after
// matching the source/destination side actions. Only one action (or
// none) will be offered here.
//
// This event can be emitted multiple times during the drag-and-drop
// operation, mainly in response to destination side changes through
// wl_data_offer.set_actions, and as the data device enters/leaves
// surfaces.
//
// It is only possible to receive this event after
// wl_data_source.dnd_drop_performed if the drag-and-drop operation
// ended in an "ask" action, in which case the final wl_data_source.action
// event will happen immediately before wl_data_source.dnd_finished.
//
// Compositors may also change the selected action on the fly, mainly
// in response to keyboard modifier changes during the drag-and-drop
// operation.
//
// The most recent action received is always the valid one. The chosen
// action may change alongside negotiation (e.g. an "ask" action can turn
// into a "move" operation), so the effects of the final action must
// always be applied in wl_data_offer.dnd_finished.
//
// Clients can trigger cursor surface changes from this point, so
// they reflect the current action.
func (o *WlDataSource) ServeAction(dndAction wire.Uint) error {
	if o.actionHandler == nil {
		return nil
	}
	return o.actionHandler(dndAction)
}

// HandleAction registers a handler for a Action event.
func (o *WlDataSource) HandleAction(h WlDataSourceActionHandler) {
	o.actionHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlDataSource) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.targetHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		mimeType, err := r.ReadString()
		if err != nil {
			return err
		}

		return o.targetHandler(mimeType)

	case 1:
		if o.sendHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		mimeType, err := r.ReadString()
		if err != nil {
			return err
		}

		fd, err := r.ReadFD()
		if err != nil {
			return err
		}

		return o.sendHandler(mimeType, fd)

	case 2:
		if o.cancelledHandler == nil {
			return nil
		}

		return o.cancelledHandler()

	case 3:
		if o.dndDropPerformedHandler == nil {
			return nil
		}

		return o.dndDropPerformedHandler()

	case 4:
		if o.dndFinishedHandler == nil {
			return nil
		}

		return o.dndFinishedHandler()

	case 5:
		if o.actionHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		dndAction, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.actionHandler(dndAction)

	default:
		return fmt.Errorf("WlDataSource: unhandled message(%v)", msg)
	}
}
