package test

import (
	"encoding/json"
	"github.com/google/uuid"
)

type FakeEvent struct{}

func NewFakeEvent() *FakeEvent {
	return &FakeEvent{}
}

func (evt *FakeEvent) Height() int64  { return 1 }
func (evt *FakeEvent) Name() string   { return "" }
func (evt *FakeEvent) Version() int   { return 0 }
func (evt *FakeEvent) Id() string     { return uuid.New().String() }
func (evt *FakeEvent) String() string {
	const FakeEventJson = `
	{
		"id": "aec351e8-9351-46b8-9ac0-f7d96199edf6",
		"height": 1,
		"name": "FakeEvent",
		"version": 1
		"payload": {
			"foo": "bar"
		}
	}
	`
	event, _ := json.Marshal(FakeEventJson)
	return string(event)
}

