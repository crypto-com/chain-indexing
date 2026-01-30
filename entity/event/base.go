package event

import (
	"github.com/google/uuid"
)

// Base is a JSON-compatible rdbprojectionbase event with sequence and EventUUID support. It is not a must to
// use this as rdbprojectionbase event but to implement your own one rdbprojectionbase on your design
type Base struct {
	EventName    string `json:"name"`
	EventVersion int    `json:"version"`
	BlockHeight  int64  `json:"height"`
	EventUUID    string `json:"uuid"`
	MsgVersion   string `json:"msgVersion"`
}

func NewBase(params BaseParams) Base {
	return Base{
		EventName:    params.Name,
		EventVersion: params.Version,
		BlockHeight:  params.BlockHeight,
		EventUUID:    uuid.New().String(),
		MsgVersion:   params.MsgVersion,
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

func (event *Base) UUID() string {
	return event.EventUUID
}

func (event *Base) APIVersion() string {
	return event.MsgVersion
}

type BaseParams struct {
	Name        string
	Version     int
	BlockHeight int64
	MsgVersion  string
}
