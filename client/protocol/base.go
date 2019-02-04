package protocol

import "github.com/kibumh/wayland-go/wire"

// Base is a base struct for all generated protocol structs.
type Base struct {
	Conn *wire.Conn
	// in     <-chan wire.Message
	// out    chan<- wire.Message
	id wire.ID
	// idchan <-chan wire.ID
}

// func NewBase(id wire.ID) Base {
// 	return Base{
// 		// in:     b.in,
// 		// out:    b.out,
// 		conn:
// 		id:     id,
// 		// idchan: b.idchan,
// 	}
// }

func (b Base) ID() wire.ID {
	return b.id
}

// func (b Base) newId() wire.ID {
// 	return <-b.idchan
// }
