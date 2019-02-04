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

// WlDataOfferError is for .
//
//
type WlDataOfferError int

const (
	WlDataOfferErrorInvalidFinish     WlDataOfferError = 0 // finish request was called untimely
	WlDataOfferErrorInvalidActionMask WlDataOfferError = 1 // action mask contains invalid values
	WlDataOfferErrorInvalidAction     WlDataOfferError = 2 // action argument has an invalid value
	WlDataOfferErrorInvalidOffer      WlDataOfferError = 3 // offer doesn't accept this request
)

/*
 * EVENT HANDLER TYPES
 */

// WlDataOfferOfferHandler is a handler for advertise offered mime type.
//
// Sent immediately after creating the wl_data_offer object.  One
// event per offered mime type.
type WlDataOfferOfferHandler func(mimeType wire.String) error

// WlDataOfferSourceActionsHandler is a handler for notify the source-side available actions.
//
// This event indicates the actions offered by the data source. It
// will be sent right after wl_data_device.enter, or anytime the source
// side changes its offered actions through wl_data_source.set_actions.
type WlDataOfferSourceActionsHandler func(sourceActions wire.Uint) error

// WlDataOfferActionHandler is a handler for notify the selected action.
//
// This event indicates the action selected by the compositor after
// matching the source/destination side actions. Only one action (or
// none) will be offered here.
//
// This event can be emitted multiple times during the drag-and-drop
// operation in response to destination side action changes through
// wl_data_offer.set_actions.
//
// This event will no longer be emitted after wl_data_device.drop
// happened on the drag-and-drop destination, the client must
// honor the last action received, or the last preferred one set
// through wl_data_offer.set_actions when handling an "ask" action.
//
// Compositors may also change the selected action on the fly, mainly
// in response to keyboard modifier changes during the drag-and-drop
// operation.
//
// The most recent action received is always the valid one. Prior to
// receiving wl_data_device.drop, the chosen action may change (e.g.
// due to keyboard modifiers being pressed). At the time of receiving
// wl_data_device.drop the drag-and-drop destination must honor the
// last action received.
//
// Action changes may still happen after wl_data_device.drop,
// especially on "ask" actions, where the drag-and-drop destination
// may choose another action afterwards. Action changes happening
// at this stage are always the result of inter-client negotiation, the
// compositor shall no longer be able to induce a different action.
//
// Upon "ask" actions, it is expected that the drag-and-drop destination
// may potentially choose a different action and/or mime type,
// based on wl_data_offer.source_actions and finally chosen by the
// user (e.g. popping up a menu with the available options). The
// final wl_data_offer.set_actions and wl_data_offer.accept requests
// must happen before the call to wl_data_offer.finish.
type WlDataOfferActionHandler func(dndAction wire.Uint) error

/*
 * TYPE
 */
// WlDataOffer is offer to transfer data.
//
// A wl_data_offer represents a piece of data offered for transfer
// by another client (the source client).  It is used by the
// copy-and-paste and drag-and-drop mechanisms.  The offer
// describes the different mime types that the data can be
// converted to and provides the mechanism for transferring the
// data directly from the source client.
type WlDataOffer struct {
	Base

	offerHandler         WlDataOfferOfferHandler
	sourceActionsHandler WlDataOfferSourceActionsHandler
	actionHandler        WlDataOfferActionHandler
}

// NewWlDataOffer creates a WlDataOffer object.
func NewWlDataOffer(c *wire.Conn) *WlDataOffer {
	return NewWlDataOfferWithID(c, c.NewID())
}

// NewWlDataOfferWithID creates a WlDataOffer object with a given id.
func NewWlDataOfferWithID(c *wire.Conn, id wire.ID) *WlDataOffer {
	o := &WlDataOffer{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// Accept is for accept one of the offered mime types
//
// Indicate that the client can accept the given mime type, or
// NULL for not accepted.
//
// For objects of version 2 or older, this request is used by the
// client to give feedback whether the client can receive the given
// mime type, or NULL if none is accepted; the feedback does not
// determine whether the drag-and-drop operation succeeds or not.
//
// For objects of version 3 or newer, this request determines the
// final result of the drag-and-drop operation. If the end result
// is that no mime types were accepted, the drag-and-drop operation
// will be cancelled and the corresponding drag source will receive
// wl_data_source.cancelled. Clients may still use this event in
// conjunction with wl_data_source.action for feedback.
func (o *WlDataOffer) Accept(serial wire.Uint, mimeType wire.String) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(serial); err != nil {
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

// Receive is for request that the data is transferred
//
// To transfer the offered data, the client issues this request
// and indicates the mime type it wants to receive.  The transfer
// happens through the passed file descriptor (typically created
// with the pipe system call).  The source client writes the data
// in the mime type representation requested and then closes the
// file descriptor.
//
// The receiving client reads from the read end of the pipe until
// EOF and then closes its end, at which point the transfer is
// complete.
//
// This request may happen multiple times for different mime types,
// both before and after wl_data_device.drop. Drag-and-drop destination
// clients may preemptively fetch data or examine it more closely to
// determine acceptance.
func (o *WlDataOffer) Receive(mimeType wire.String, fd wire.FD) error {
	msg, err := wire.NewMessage(o.ID(), 1)
	if err != nil {
		return err
	}

	if err = msg.Write(mimeType); err != nil {
		return err
	}

	if err = msg.Write(fd); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Destroy is for destroy data offer
//
// Destroy the data offer.
func (o *WlDataOffer) Destroy() error {
	msg, err := wire.NewMessage(o.ID(), 2)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// Finish is for the offer will no longer be used
//
// Notifies the compositor that the drag destination successfully
// finished the drag-and-drop operation.
//
// Upon receiving this request, the compositor will emit
// wl_data_source.dnd_finished on the drag source client.
//
// It is a client error to perform other requests than
// wl_data_offer.destroy after this one. It is also an error to perform
// this request after a NULL mime type has been set in
// wl_data_offer.accept or no action was received through
// wl_data_offer.action.
func (o *WlDataOffer) Finish() error {
	msg, err := wire.NewMessage(o.ID(), 3)
	if err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

// SetActions is for set the available/preferred drag-and-drop actions
//
// Sets the actions that the destination side client supports for
// this operation. This request may trigger the emission of
// wl_data_source.action and wl_data_offer.action events if the compositor
// needs to change the selected action.
//
// This request can be called multiple times throughout the
// drag-and-drop operation, typically in response to wl_data_device.enter
// or wl_data_device.motion events.
//
// This request determines the final result of the drag-and-drop
// operation. If the end result is that no action is accepted,
// the drag source will receive wl_drag_source.cancelled.
//
// The dnd_actions argument must contain only values expressed in the
// wl_data_device_manager.dnd_actions enum, and the preferred_action
// argument must only contain one of those values set, otherwise it
// will result in a protocol error.
//
// While managing an "ask" action, the destination drag-and-drop client
// may perform further wl_data_offer.receive requests, and is expected
// to perform one last wl_data_offer.set_actions request with a preferred
// action other than "ask" (and optionally wl_data_offer.accept) before
// requesting wl_data_offer.finish, in order to convey the action selected
// by the user. If the preferred action is not in the
// wl_data_offer.source_actions mask, an error will be raised.
//
// If the "ask" action is dismissed (e.g. user cancellation), the client
// is expected to perform wl_data_offer.destroy right away.
//
// This request can only be made on drag-and-drop offers, a protocol error
// will be raised otherwise.
func (o *WlDataOffer) SetActions(dndActions wire.Uint, preferredAction wire.Uint) error {
	msg, err := wire.NewMessage(o.ID(), 4)
	if err != nil {
		return err
	}

	if err = msg.Write(dndActions); err != nil {
		return err
	}

	if err = msg.Write(preferredAction); err != nil {
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

// ServeOffer is for advertise offered mime type.
//
// Sent immediately after creating the wl_data_offer object.  One
// event per offered mime type.
func (o *WlDataOffer) ServeOffer(mimeType wire.String) error {
	if o.offerHandler == nil {
		return nil
	}
	return o.offerHandler(mimeType)
}

// HandleOffer registers a handler for a Offer event.
func (o *WlDataOffer) HandleOffer(h WlDataOfferOfferHandler) {
	o.offerHandler = h
}

// ServeSourceActions is for notify the source-side available actions.
//
// This event indicates the actions offered by the data source. It
// will be sent right after wl_data_device.enter, or anytime the source
// side changes its offered actions through wl_data_source.set_actions.
func (o *WlDataOffer) ServeSourceActions(sourceActions wire.Uint) error {
	if o.sourceActionsHandler == nil {
		return nil
	}
	return o.sourceActionsHandler(sourceActions)
}

// HandleSourceActions registers a handler for a SourceActions event.
func (o *WlDataOffer) HandleSourceActions(h WlDataOfferSourceActionsHandler) {
	o.sourceActionsHandler = h
}

// ServeAction is for notify the selected action.
//
// This event indicates the action selected by the compositor after
// matching the source/destination side actions. Only one action (or
// none) will be offered here.
//
// This event can be emitted multiple times during the drag-and-drop
// operation in response to destination side action changes through
// wl_data_offer.set_actions.
//
// This event will no longer be emitted after wl_data_device.drop
// happened on the drag-and-drop destination, the client must
// honor the last action received, or the last preferred one set
// through wl_data_offer.set_actions when handling an "ask" action.
//
// Compositors may also change the selected action on the fly, mainly
// in response to keyboard modifier changes during the drag-and-drop
// operation.
//
// The most recent action received is always the valid one. Prior to
// receiving wl_data_device.drop, the chosen action may change (e.g.
// due to keyboard modifiers being pressed). At the time of receiving
// wl_data_device.drop the drag-and-drop destination must honor the
// last action received.
//
// Action changes may still happen after wl_data_device.drop,
// especially on "ask" actions, where the drag-and-drop destination
// may choose another action afterwards. Action changes happening
// at this stage are always the result of inter-client negotiation, the
// compositor shall no longer be able to induce a different action.
//
// Upon "ask" actions, it is expected that the drag-and-drop destination
// may potentially choose a different action and/or mime type,
// based on wl_data_offer.source_actions and finally chosen by the
// user (e.g. popping up a menu with the available options). The
// final wl_data_offer.set_actions and wl_data_offer.accept requests
// must happen before the call to wl_data_offer.finish.
func (o *WlDataOffer) ServeAction(dndAction wire.Uint) error {
	if o.actionHandler == nil {
		return nil
	}
	return o.actionHandler(dndAction)
}

// HandleAction registers a handler for a Action event.
func (o *WlDataOffer) HandleAction(h WlDataOfferActionHandler) {
	o.actionHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlDataOffer) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.offerHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		mimeType, err := r.ReadString()
		if err != nil {
			return err
		}

		return o.offerHandler(mimeType)

	case 1:
		if o.sourceActionsHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		sourceActions, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.sourceActionsHandler(sourceActions)

	case 2:
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
		return fmt.Errorf("WlDataOffer: unhandled message(%v)", msg)
	}
}
