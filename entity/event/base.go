package event

import (
	"github.com/google/uuid"
)

// Base is a JSON-compatible rdbbase event with sequence and UUID support. It is not a must to
// use this as rdbbase event but to implement your own one rdbbase on your design
type Base struct {
	EventName    string `json:"name"`
	EventVersion int    `json:"version"`
	BlockHeight  int64  `json:"height"`
	UUID         string `json:"id"`
}

func NewBase(params BaseParams) Base {
	return Base{
		EventName:    params.Name,
		EventVersion: params.Version,
		BlockHeight:  params.BlockHeight,
		UUID:         uuid.New().String(),
	}
}

func (event *Base) Name() string {
	return event.EventName
}

func (event *Base) Version() int {
	return event.EventVersion
}

func (event *Base) Height() int64 {
	return event.BlockHeight
}

func (event *Base) Id() string {
	return event.UUID
}

type BaseParams struct {
	Name        string
	Version     int
	BlockHeight int64
}
