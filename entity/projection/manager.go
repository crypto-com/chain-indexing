package projection

import (
	"fmt"
	"time"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/metric/prometheus"
)

// StoreBasedManager is a projection manager relies on replaying events from EventStore
type StoreBasedManager struct {
	logger     applogger.Logger
	eventStore entity_event.Store

	projections []Projection
}

func NewStoreBasedManager(logger applogger.Logger, eventStore entity_event.Store) *StoreBasedManager {
	return &StoreBasedManager{
		logger: logger.WithFields(applogger.LogFields{
			"module": "projectionManager",
		}),
		eventStore: eventStore,

		projections: make([]Projection, 0),
	}
}

func (manager *StoreBasedManager) RegisterProjection(projection Projection) error {
	if manager.IsProjectionRegistered(projection) {
		return fmt.Errorf("projection `%s` already registered", projection.Id())
	}
	manager.projections = append(manager.projections, projection)
	return nil
}

func (manager *StoreBasedManager) IsProjectionRegistered(projection Projection) bool {
	for _, registeredProjection := range manager.projections {
		if projection.Id() == registeredProjection.Id() {
			return true
		}
	}
	return false
}

// Starts projectionManager by running all registered projection.
func (manager *StoreBasedManager) RunInBackground() {
	for _, projection := range manager.projections {
		go manager.projectionRunner(projection)
	}
}

func (manager *StoreBasedManager) projectionRunner(projection Projection) {
	eventsToListen := projection.GetEventsToListen()
	logger := manager.logger.WithFields(applogger.LogFields{
		"projection": projection.Id(),
	})

	logger.WithFields(applogger.LogFields{
		"eventsToListen": eventsToListen,
	}).Infof("projection start running")

	var lastHandledEventHeight *int64
	for {
		var err error
		lastHandledEventHeight, err = projection.GetLastHandledEventHeight()
		if err == nil {
			break
		}

		logger.Infof("error getting last handled event height from projection")
		<-waitFor(5 * time.Second)
	}

	var nextEventHeight int64
	if lastHandledEventHeight == nil {
		nextEventHeight = 0
	} else {
		nextEventHeight = *lastHandledEventHeight + 1
	}

	for {
		latestEventHeight, _ := manager.eventStore.GetLatestHeight()
		if latestEventHeight == nil {
			logger.Debugf("no event in in the system yet")
			<-waitFor(5 * time.Second)
			continue
		}
		for nextEventHeight <= *latestEventHeight {
			startTime := time.Now()
			var err error

			eventLogger := logger.WithFields(applogger.LogFields{
				"height": nextEventHeight,
			})

			var eventsAtHeight []entity_event.Event
			if eventsAtHeight, err = manager.eventStore.GetAllByHeight(nextEventHeight); err != nil {
				eventLogger.Errorf("error getting all events by height: %v", err)
				<-waitFor(time.Second)
				continue
			}

			var events = make([]entity_event.Event, 0)
			for _, event := range eventsAtHeight {
				if !isListeningEvent(event, eventsToListen) {
					//eventLogger.WithFields(applogger.LogFields{
					//	"eventName": event.Name(),
					//}).Debugf("skipping because event is not one of the listening events")
					continue
				}
				events = append(events, event)
			}

			eventLogger = eventLogger.WithFields(applogger.LogFields{
				"eventCount": len(events),
			})
			if err = projection.HandleEvents(nextEventHeight, events); err != nil {
				eventLogger.WithFields(applogger.LogFields{
					"events": events,
				}).Errorf("error handling events: %v", err)
				<-waitFor(5 * time.Second)
				continue
			}

			eventLogger.Infof("successfully handled events")
			prometheus.RecordProjectionExecTime(projection.Id(), time.Since(startTime).Milliseconds())
			nextEventHeight += 1
		}
		prometheus.RecordProjectionLatestHeight(projection.Id(), nextEventHeight)
		<-waitFor(5 * time.Second)
	}
}

func isListeningEvent(event entity_event.Event, eventsToListen []string) bool {
	targetEventName := event.Name()
	for _, eventName := range eventsToListen {
		if targetEventName == eventName {
			return true
		}
	}

	return false
}

func waitFor(wait time.Duration) <-chan time.Time {
	return time.After(wait)
}
