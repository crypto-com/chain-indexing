package projection_test

import (
	"time"

	. "github.com/crypto-com/chainindex/entity/event/test"
	. "github.com/crypto-com/chainindex/entity/projection/test"
	. "github.com/crypto-com/chainindex/internal/logger/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/entity/projection"
	"github.com/crypto-com/chainindex/internal/primptr"
)

var _ = Describe("Manager", func() {
	Describe("RegisterProjection", func() {
		It("should return Error when the projection with the same Id already exists", func() {
			manager := projection.NewManager(NewFakeLogger(), NewFakeEventStore())

			anyProjection := NewFakeProjection()
			Expect(manager.RegisterProjection(anyProjection)).To(BeNil())
			Expect(manager.RegisterProjection(anyProjection)).To(MatchError("projection `FakeProjection` already registered"))
		})

		It("should register projection to the manager", func() {
			manager := projection.NewManager(NewFakeLogger(), NewFakeEventStore())

			anyProjection := NewFakeProjection()

			Expect(manager.IsProjectionRegistered(anyProjection)).To(BeFalse())

			Expect(manager.RegisterProjection(anyProjection)).To(BeNil())

			Expect(manager.IsProjectionRegistered(anyProjection)).To(BeTrue())
		})

		It("should pass only listening events to projection", func() {
			// Setup
			mockEventStore := NewMockEventStore()
			manager := projection.NewManager(NewFakeLogger(), mockEventStore)
			mockProjection := NewMockProjection()

			// BlockEvent setup
			anyEvent := newAnyEvent()
			anyOtherEvent := newAnyOtherEvent()

			// Projection setup
			anyProjectionId := "ANY_PROJECTION_ID"
			mockProjection.On("Id").Return(anyProjectionId)
			mockProjection.On("GetEventsToListen").Return([]string{anyEvent.Name()})
			mockProjection.On("GetLastHandledEventHeight").Return(
				primptr.Int64(0), nil,
			)

			// Register the Projection
			err := manager.RegisterProjection(mockProjection)
			Expect(err).To(BeNil())

			// Produce event to the event store
			nextHeight := int64(1)
			mockEventStore.On("GetAllByHeight", nextHeight).Return(
				[]entity_event.Event{anyEvent, anyOtherEvent}, nil,
			)
			mockEventStore.On("GetLatestHeight").Return(primptr.Int64(int64(1)), nil)

			// Define the assertion expectations
			mockProjection.On("HandleEvents", nextHeight, mock.MatchedBy(func(events interface{}) bool {
				typedEvents, _ := events.([]entity_event.Event)
				// should only receive `anyEvent`
				return typedEvents[0].Name() == anyEvent.Name()
			})).Once().Return(nil)

			// Run the manager
			manager.Run()
			// Since manager create goroutines for Projection, we have to give up the CPU for
			// events channel to happen
			<-time.After(time.Second)

			// Assert the projection expectations. i.e. events are handled
			mockProjection.AssertExpectations(GinkgoT())
		})

		It("should pass all the unhandled events to projection", func() {
			// Setup
			mockEventStore := NewMockEventStore()
			manager := projection.NewManager(NewFakeLogger(), mockEventStore)
			mockProjection := NewMockProjection()

			// BlockEvent setup
			anyEvent := newAnyEvent()
			anyOtherEvent := newAnyOtherEvent()

			// Projection setup
			anyProjectionId := "ANY_PROJECTION_ID"
			mockProjection.On("Id").Return(anyProjectionId)
			mockProjection.On("GetEventsToListen").Return([]string{anyEvent.Name(), anyOtherEvent.Name()})
			mockProjection.On("GetLastHandledEventHeight").Return(
				primptr.Int64(0), nil,
			)

			// Register the Projection
			err := manager.RegisterProjection(mockProjection)
			Expect(err).To(BeNil())

			// Produce event to the event store
			anyEventHeight := int64(1)
			mockEventStore.On("GetAllByHeight", anyEventHeight).Return(
				[]entity_event.Event{anyEvent}, nil,
			)
			anyOtherEventHeight := int64(2)
			mockEventStore.On("GetAllByHeight", anyOtherEventHeight).Return(
				[]entity_event.Event{anyOtherEvent}, nil,
			)
			mockEventStore.On("GetLatestHeight").Return(primptr.Int64(int64(2)), nil)

			// Define the assertion expectations
			mockProjection.On("HandleEvents", anyEventHeight, mock.MatchedBy(func(events interface{}) bool {
				typedEvents, _ := events.([]entity_event.Event)
				return len(typedEvents) == 1 && typedEvents[0].Name() == anyEvent.Name()
			})).Once().Return(nil)
			mockProjection.On("HandleEvents", anyOtherEventHeight, mock.MatchedBy(func(events interface{}) bool {
				typedEvents, _ := events.([]entity_event.Event)
				return len(typedEvents) == 1 && typedEvents[0].Name() == anyOtherEvent.Name()
			})).Once().Return(nil)

			// Run the manager
			manager.Run()
			// Since manager create goroutines for Projection, we have to give up the CPU for
			// events channel to happen
			<-time.After(time.Second)

			// Assert the projection expectations. i.e. events are handled
			mockProjection.AssertExpectations(GinkgoT())
		})

		It("should work for multiple projections", func() {
			var err error

			// Setup
			mockEventStore := NewMockEventStore()
			manager := projection.NewManager(NewFakeLogger(), mockEventStore)
			anyProjection := NewMockProjection()
			anyOtherProjection := NewMockProjection()

			// BlockEvent setup
			anyEvent := newAnyEvent()
			anyOtherEvent := newAnyOtherEvent()

			// Projection setup
			anyProjectionId := "ANY_PROJECTION_ID"
			anyProjection.On("Id").Return(anyProjectionId)
			anyProjection.On("GetEventsToListen").Return([]string{anyEvent.Name()})
			anyProjection.On("GetLastHandledEventHeight").Return(
				primptr.Int64(0), nil,
			)

			anyOtherProjectionId := "ANY_OTHER_PROJECTION_ID"
			anyOtherProjection.On("Id").Return(anyOtherProjectionId)
			anyOtherProjection.On("GetEventsToListen").Return([]string{anyOtherEvent.Name()})
			anyOtherProjection.On("GetLastHandledEventHeight").Return(
				primptr.Int64(0), nil,
			)

			// Register the Projection
			err = manager.RegisterProjection(anyProjection)
			Expect(err).To(BeNil())
			err = manager.RegisterProjection(anyOtherProjection)
			Expect(err).To(BeNil())

			// Produce event to the event store
			nextHeight := int64(1)
			mockEventStore.On("GetAllByHeight", nextHeight).Return(
				[]entity_event.Event{anyEvent, anyOtherEvent}, nil,
			)
			mockEventStore.On("GetLatestHeight").Return(primptr.Int64(nextHeight), nil)

			// Define the assertion expectations
			anyProjection.On("HandleEvents", nextHeight, mock.MatchedBy(func(events interface{}) bool {
				typedEvents, _ := events.([]entity_event.Event)
				return len(typedEvents) == 1 && typedEvents[0].Name() == anyEvent.Name()
			})).Once().Return(nil)
			anyOtherProjection.On("HandleEvents", nextHeight, mock.MatchedBy(func(events interface{}) bool {
				typedEvents, _ := events.([]entity_event.Event)
				return len(typedEvents) == 1 && typedEvents[0].Name() == anyOtherEvent.Name()
			})).Once().Return(nil)

			// Run the manager
			manager.Run()
			// Since manager create goroutines for Projection, we have to give up the CPU for
			// events channel to happen
			<-time.After(time.Second)

			// Assert the projection expectations. i.e. events are handled
			anyProjection.AssertExpectations(GinkgoT())
			anyOtherProjection.AssertExpectations(GinkgoT())
		})

		It("should pass event to projection based on last handled event height", func() {
			// Setup
			mockEventStore := NewMockEventStore()
			manager := projection.NewManager(NewFakeLogger(), mockEventStore)
			mockProjection := NewMockProjection()

			// BlockEvent setup
			anyEvent := newAnyEvent()

			// Projection setup
			anyProjectionId := "ANY_PROJECTION_ID"
			mockProjection.On("Id").Return(anyProjectionId)
			mockProjection.On("GetEventsToListen").Return([]string{anyEvent.Name()})
			mockProjection.On("GetLastHandledEventHeight").Return(
				primptr.Int64(int64(1)), nil,
			)

			// Register the Projection
			err := manager.RegisterProjection(mockProjection)
			Expect(err).To(BeNil())

			// Produce event to the event store
			nextHeight := int64(2)
			mockEventStore.On("GetAllByHeight", nextHeight).Return(
				[]entity_event.Event{anyEvent}, nil,
			)
			mockEventStore.On("GetLatestHeight").Return(primptr.Int64(nextHeight), nil)

			// Define the assertion expectations
			mockProjection.On("HandleEvents", nextHeight, mock.MatchedBy(func(events interface{}) bool {
				typedEvents, _ := events.([]entity_event.Event)
				return len(typedEvents) == 1 && typedEvents[0].Name() == anyEvent.Name()
			})).Once().Return(nil)

			// Run the manager
			manager.Run()
			// Since manager create goroutines for Projection, we have to give up the CPU for
			// events channel to happen
			<-time.After(time.Second)

			// Assert the projection expectations. i.e. events are handled
			mockProjection.AssertExpectations(GinkgoT())
		})
	})
})

func newAnyEvent() entity_event.Event {
	anyEventName := "ANY_EVENT"
	anyEvent := NewMockEvent()
	anyEvent.On("Name").Return(anyEventName)

	return anyEvent
}

func newAnyOtherEvent() entity_event.Event {
	anyOtherEventName := "ANY_OTHER_EVENT"
	anyOtherEvent := NewMockEvent()
	anyOtherEvent.On("Name").Return(anyOtherEventName)

	return anyOtherEvent
}
