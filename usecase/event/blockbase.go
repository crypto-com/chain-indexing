package event

import "github.com/gofrs/uuid"

// Base is a JSON-compatible base event with sequence and UUID support. It is not a must to
// use this as base event but to implement your own one base on your design
type BlockBase struct {
	Height int64  `json:"height"`
	UUID   string `json:"uuid"`
}

func NewBlockBase(height int64) BlockBase {
	return BlockBase{
		Height: height,
	}
}

func (evt BlockBase) Id() string {
	// Lazy creation of UUID
	if evt.UUID == "" {
		evt.UUID = uuid.Must(uuid.NewV4()).String()
	}
	return evt.UUID
}
