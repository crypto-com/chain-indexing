package event_test

import (
	"bytes"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/event"
)

var _ = Describe("eventRegistry", func() {
	Describe("Register", func() {
		It("should register an event decoder with the name and version", func() {
			registry := event.NewRegistry()

			Expect(registry.IsRegistered(simpleJSONEventName, simpleJSONEventVersion)).To(Equal(false))

			registry.Register(simpleJSONEventName, simpleJSONEventVersion, decodeSimpleJSONEvent)

			Expect(registry.IsRegistered(simpleJSONEventName, simpleJSONEventVersion)).To(Equal(true))
		})
	})

	Describe("DecodeByType", func() {
		It("should return error when the event type is not registered", func() {
			var err error

			registry := event.NewRegistry()
			anyEvent := newSimpleJSONEvent()
			encoded, _ := json.Marshal(anyEvent)

			_, err = registry.DecodeByType(anyEvent.Name(), anyEvent.Version(), encoded)
			Expect(err).To(MatchError("unrecognized event type `SimpleJSONEventV0`"))
		})

		It("should return error when the encoded event does not match with the provided type", func() {
			var err error

			registry := event.NewRegistry()
			registry.Register(simpleJSONEventName, simpleJSONEventVersion, decodeSimpleJSONEvent)
			anyMismatchEncoded := []byte("{\"invalid\":true}")

			_, err = registry.DecodeByType(simpleJSONEventName, simpleJSONEventVersion, anyMismatchEncoded)
			Expect(err).To(MatchError("error decoding event: json: unknown field \"invalid\""))
		})

		It("should return the decoded event based on type", func() {
			var err error

			registry := event.NewRegistry()
			registry.Register(simpleJSONEventName, simpleJSONEventVersion, decodeSimpleJSONEvent)
			anyEvent := newSimpleJSONEvent()
			encoded, _ := json.Marshal(anyEvent)

			var actual event.Event
			actual, err = registry.DecodeByType(anyEvent.Name(), anyEvent.Version(), encoded)
			Expect(err).To(BeNil())
			typedEvent, ok := actual.(*simpleJSONEvent)
			Expect(ok).To(BeTrue())
			Expect(typedEvent).To(Equal(newSimpleJSONEvent()))
		})
	})
})

const simpleJSONEventName = "SimpleJSONEvent"
const simpleJSONEventVersion = 0

type simpleJSONEvent struct {
	Key string `json:"key"`
}

func newSimpleJSONEvent() *simpleJSONEvent {
	return &simpleJSONEvent{Key: "value"}
}
func (event *simpleJSONEvent) Height() int64 { return 1 }
func (event *simpleJSONEvent) Name() string  { return simpleJSONEventName }
func (event *simpleJSONEvent) Version() int  { return simpleJSONEventVersion }
func (event *simpleJSONEvent) UUID() string  { return "simple-json-event-id" }
func (event *simpleJSONEvent) ToJSON() (string, error) {
	return event.UUID(), nil
}
func (event *simpleJSONEvent) String() string { return simpleJSONEventName }
func decodeSimpleJSONEvent(eventBytes []byte) (event.Event, error) {
	var event *simpleJSONEvent
	jsonDecoder := json.NewDecoder(bytes.NewReader(eventBytes))
	jsonDecoder.DisallowUnknownFields()
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
