package wire

type Int uint32
type Uint uint32
type Fixed float64
type String string
type ID uint32
type Object interface {
	ID() ID
	ServeMessage(*Message) error
}
type NewID uint32
type Array []byte
type FD uintptr
