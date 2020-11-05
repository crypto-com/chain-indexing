package projection_test

// TODO
// import (
// 	"time"

// 	"github.com/onsi/ginkgo"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/stretchr/testify/mock"

// 	ddd_event_test "github.com/crypto-com/chainindex/ddd/event/test"
// 	"github.com/crypto-com/chainindex/ddd/projection"
// 	. "github.com/crypto-com/chainindex/ddd/projection/test"
// 	"github.com/crypto-com/chainindex/internal/primptr"
// 	. "github.com/crypto-com/chainindex/test/fake"
// )

// var _ = Describe("ProjectionManager", func() {
// 	Describe("RegisterProjection", func() {
// 		It("should return Error when the projection with the same Id already exists", func() {
// 			manager := projection.NewManager(NewFakeLogger(), ddd_event_test.NewFakeManager())

// 			anyProjection := NewFakeProjection()
// 			Expect(manager.RegisterProjection(anyProjection)).To(BeNil())
// 			Expect(manager.RegisterProjection(anyProjection)).To(MatchError("Projection `FakeProjection` already registered"))
// 		})

// 		It("should register projection to the manager", func() {
// 			manager := projection.NewManager(NewFakeLogger(), ddd_event_test.NewFakeManager())

// 			anyProjection := NewFakeProjection()

// 			Expect(manager.IsProjectionRegistered(anyProjection)).To(BeFalse())

// 			Expect(manager.RegisterProjection(anyProjection)).To(BeNil())

// 			Expect(manager.IsProjectionRegistered(anyProjection)).To(BeTrue())
// 		})

// 		It("should pass all the unhandled events to projection", func() {
// 			// Setup
// 			mockEventManager := ddd_event_test.NewMockManager()
// 			manager := projection.NewManager(NewFakeLogger(), mockEventManager)
// 			mockProjection := NewMockProjection()

// 			// Event setup
// 			anyEventName := "ANY_EVENT"
// 			anyEvent := ddd_event_test.NewMockEvent()
// 			anyEvent.On("Name").Return(anyEventName)
// 			anyOtherEventName := "ANY_OTHER_EVENT"
// 			anyOtherEvent := ddd_event_test.NewMockEvent()
// 			anyOtherEvent.On("Name").Return(anyOtherEventName)

// 			// Projection setup
// 			anyProjectionId := "ANY_PROJECTION_ID"
// 			mockProjection.On("Id").Return(anyProjectionId)
// 			mockProjection.On("GetEventsToListen").Return([]string{anyEventName, anyOtherEventName})
// 			mockProjection.On("GetLastHandledEventSeq").Return(primptr.Int64Nil())

// 			// Register the Projection
// 			err := manager.RegisterProjection(mockProjection)
// 			Expect(err).To(BeNil())

// 			// Produce event to the event manager
// 			mockEventManager.On("GetEventBySeq", int64(0)).Return(anyEvent, nil)
// 			mockEventManager.On("GetEventBySeq", int64(1)).Return(anyOtherEvent, nil)
// 			mockEventManager.On("GetLatestEventSeq").Return(primptr.Int64(int64(1)), nil)

// 			// Define the assertion expectations
// 			mockProjection.On("HandleEvent", mock.MatchedBy(func(evt interface{}) bool {
// 				return evt.(*ddd_event_test.MockEvent).Name() == anyEvent.Name()
// 			})).Once().Return(nil)
// 			mockProjection.On("HandleEvent", mock.MatchedBy(func(evt interface{}) bool {
// 				return evt.(*ddd_event_test.MockEvent).Name() == anyOtherEvent.Name()
// 			})).Once().Return(nil)

// 			// Run the manager
// 			manager.Run()
// 			// Since manager create goroutines for Projection, we have to give up the CPU for
// 			// events channel to happen
// 			<-time.After(time.Second)

// 			// Assert the projection expectations. i.e. events are handled
// 			mockProjection.AssertExpectations(ginkgo.GinkgoT())
// 		})

// 		It("should work for multiple projections", func() {
// 			var err error

// 			// Setup
// 			mockEventManager := ddd_event_test.NewMockManager()
// 			manager := projection.NewManager(NewFakeLogger(), mockEventManager)
// 			anyProjection := NewMockProjection()
// 			anyOtherProjection := NewMockProjection()

// 			// Event setup
// 			anyEventName := "ANY_EVENT"
// 			anyEvent := ddd_event_test.NewMockEvent()
// 			anyEvent.On("Name").Return(anyEventName)
// 			anyOtherEventName := "ANY_OTHER_EVENT"
// 			anyOtherEvent := ddd_event_test.NewMockEvent()
// 			anyOtherEvent.On("Name").Return(anyOtherEventName)

// 			// Projection setup
// 			anyProjectionId := "ANY_PROJECTION_ID"
// 			anyProjection.On("Id").Return(anyProjectionId)
// 			anyProjection.On("GetEventsToListen").Return([]string{anyEventName})
// 			anyProjection.On("GetLastHandledEventSeq").Return(primptr.Int64Nil())

// 			anyOtherProjectionId := "ANY_OTHER_PROJECTION_ID"
// 			anyOtherProjection.On("Id").Return(anyOtherProjectionId)
// 			anyOtherProjection.On("GetEventsToListen").Return([]string{anyOtherEventName})
// 			anyOtherProjection.On("GetLastHandledEventSeq").Return(primptr.Int64Nil())

// 			// Register the Projection
// 			err = manager.RegisterProjection(anyProjection)
// 			Expect(err).To(BeNil())
// 			err = manager.RegisterProjection(anyOtherProjection)
// 			Expect(err).To(BeNil())

// 			// Produce event to the event manager
// 			mockEventManager.On("GetEventBySeq", int64(0)).Return(anyEvent, nil)
// 			mockEventManager.On("GetEventBySeq", int64(1)).Return(anyOtherEvent, nil)
// 			mockEventManager.On("GetLatestEventSeq").Return(primptr.Int64(int64(1)), nil)

// 			// Define the assertion expectations
// 			anyProjection.On("HandleEvent", mock.MatchedBy(func(evt interface{}) bool {
// 				return evt.(*ddd_event_test.MockEvent).Name() == anyEvent.Name()
// 			})).Once().Return(nil)
// 			anyOtherProjection.On("HandleEvent", mock.MatchedBy(func(evt interface{}) bool {
// 				return evt.(*ddd_event_test.MockEvent).Name() == anyOtherEvent.Name()
// 			})).Once().Return(nil)

// 			// Run the manager
// 			manager.Run()
// 			// Since manager create goroutines for Projection, we have to give up the CPU for
// 			// events channel to happen
// 			<-time.After(time.Second)

// 			// Assert the projection expectations. i.e. events are handled
// 			anyProjection.AssertExpectations(ginkgo.GinkgoT())
// 			anyOtherProjection.AssertExpectations(ginkgo.GinkgoT())
// 		})

// 		It("should pass event to projection based on last handled event sequence", func() {
// 			// Setup
// 			mockEventManager := ddd_event_test.NewMockManager()
// 			manager := projection.NewManager(NewFakeLogger(), mockEventManager)
// 			mockProjection := NewMockProjection()

// 			// Event setup
// 			anyEventName := "ANY_EVENT"
// 			anyOtherEvent := ddd_event_test.NewMockEvent()
// 			anyOtherEvent.On("Name").Return(anyEventName)

// 			// Projection setup
// 			anyProjectionId := "ANY_PROJECTION_ID"
// 			mockProjection.On("Id").Return(anyProjectionId)
// 			mockProjection.On("GetEventsToListen").Return([]string{anyEventName})
// 			mockProjection.On("GetLastHandledEventSeq").Return(primptr.Int64(int64(0)))

// 			// Register the Projection
// 			err := manager.RegisterProjection(mockProjection)
// 			Expect(err).To(BeNil())

// 			// Produce event to the event manager
// 			mockEventManager.On("GetEventBySeq", int64(1)).Return(anyOtherEvent, nil)
// 			mockEventManager.On("GetLatestEventSeq").Return(primptr.Int64(int64(1)), nil)

// 			// Define the assertion expectations
// 			mockProjection.On("HandleEvent", mock.MatchedBy(func(evt interface{}) bool {
// 				return evt.(*ddd_event_test.MockEvent).Name() == anyOtherEvent.Name()
// 			})).Once().Return(nil)

// 			// Run the manager
// 			manager.Run()
// 			// Since manager create goroutines for Projection, we have to give up the CPU for
// 			// events channel to happen
// 			<-time.After(time.Second)

// 			// Assert the projection expectations. i.e. events are handled
// 			mockProjection.AssertExpectations(ginkgo.GinkgoT())
// 		})
// 	})
// })
