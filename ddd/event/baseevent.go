package event

import "github.com/gofrs/uuid"

// Base is a JSON-compatible base event with sequence and UUID support. It is not a must to
// use this as base event but to implement your own one base on your design
type Base struct {
	Sequence *int64 `json:"sequence"`
	UUID     string `json:"uuid"`
}

func NewBase() Base {
	return Base{}
}

func (evt Base) MaybeSeq() *int64 {
	return evt.Sequence
}

func (evt Base) SetSeq(seq int64) {
	evt.Sequence = &seq
}

func (evt Base) Id() string {
	// Lazy creation of UUID
	if evt.UUID == "" {
		evt.UUID = uuid.Must(uuid.NewV4()).String()
	}
	return evt.UUID
}
