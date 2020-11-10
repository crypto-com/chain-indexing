package event

import (
	"errors"
	"fmt"
)

var ErrMismatchEvent = errors.New("mismatched event")

type Registry struct {
	decoders map[string]Decoder
}

func NewRegistry() *Registry {
	return &Registry{
		decoders: make(map[string]Decoder),
	}
}

// Register add a mapping of "event name ane version" to Decoder to the registry. It will overwrite
// existing registration if any.
func (registry *Registry) Register(eventName string, eventVersion int, decoder Decoder) {
	registry.decoders[eventType(eventName, eventVersion)] = decoder
}

// IsRegister returns true when the event to decoder mapping is already registered
func (registry *Registry) IsRegistered(eventName string, eventVersion int) bool {
	_, exist := registry.decoders[eventType(eventName, eventVersion)]
	return exist
}

func (registry *Registry) DecodeByType(eventName string, eventVersion int, encoded []byte) (Event, error) {
	var err error

	if !registry.IsRegistered(eventName, eventVersion) {
		return nil, fmt.Errorf("unrecognized event type `%s`", eventType(eventName, eventVersion))
	}

	decoder := registry.decoders[eventType(eventName, eventVersion)]
	var event Event
	if event, err = decoder(encoded); err != nil {
		return nil, fmt.Errorf("error decoding event: %v", err)
	}

	return event, nil
}

func eventType(name string, version int) string {
	return fmt.Sprintf("%sV%d", name, version)
}

type Decoder = func([]byte) (Event, error)
