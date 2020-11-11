package event

import "github.com/google/uuid"

// Base is a JSON-compatible rdbbase event with sequence and UUID support. It is not a must to
// use this as rdbbase event but to implement your own one rdbbase on your design
type Base struct {
	BlockHeight int64  `json:"height"`
	UUID        string `json:"id"`
}

func NewBase(height int64) Base {
	return Base{
		BlockHeight: height,
		UUID:        uuid.New().String(),
	}
}

func (evt Base) Height() int64 {
	return evt.BlockHeight
}

func (evt Base) Id() string {
	return evt.UUID
}
