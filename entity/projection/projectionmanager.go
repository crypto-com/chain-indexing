package projection

import (
	"fmt"
	"time"

	"github.com/crypto-com/chainindex/ddd/event"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type Manager struct {
	logger       applogger.Logger
	eventManager event.Manager

	projections []Projection
}

func NewManager(logger applogger.Logger, eventManager event.Manager) *Manager {
	return &Manager{
		logger:       logger,
		eventManager: eventManager,

		projections: make([]Projection, 0),
	}
}

func (manager *Manager) RegisterProjection(projection Projection) error {
	if manager.IsProjectionRegistered(projection) {
		return fmt.Errorf("Projection `%s` already registered", projection.Id())
	}
	manager.projections = append(manager.projections, projection)
	return nil
}

func (manager *Manager) IsProjectionRegistered(projection Projection) bool {
	for _, registeredProjection := range manager.projections {
		if projection.Id() == registeredProjection.Id() {
			return true
		}
	}
	return false
}

// Starts ProjectionManager by running all registered projection.
func (manager *Manager) Run() {
	for _, projection := range manager.projections {
		go manager.projectionRunner(projection)
	}
}

func (manager *Manager) projectionRunner(projection Projection) {
	evtsToListen := projection.GetEventsToListen()
	lastHandledEvtSeq := projection.GetLastHandledEventSeq()
	logger := manager.logger.WithFields(applogger.LogFields{
		"projection":     projection.Id(),
		"eventsToListen": evtsToListen,
	})

	var nextEvtSeqToHandle int64
	if lastHandledEvtSeq == nil {
		nextEvtSeqToHandle = 0
	} else {
		nextEvtSeqToHandle = *lastHandledEvtSeq + 1
	}

	for {
		latestEvtSeq := manager.eventManager.GetLatestEventSeq()
		if latestEvtSeq == nil {
			logger.Debugf("No event in in the system yet")
			<-waitToRetry(5 * time.Second)
		}
		for ; nextEvtSeqToHandle <= *latestEvtSeq; nextEvtSeqToHandle += 1 {
			var err error

			eventLogger := logger.WithFields(applogger.LogFields{
				"sequence": nextEvtSeqToHandle,
			})

			var evt event.Event
			if evt, err = manager.eventManager.GetEventBySeq(nextEvtSeqToHandle); err != nil {
				eventLogger.Errorf("Error getting event by sequence: %v", err)
				<-waitToRetry(time.Second)
				continue
			}

			eventLogger = eventLogger.WithFields(applogger.LogFields{
				"event": evt,
			})
			if !isListeningEvent(evt, evtsToListen) {
				eventLogger.Debugf("Skipping because event is not one of the listening events")
				continue
			}

			if err = projection.HandleEvent(evt); err != nil {
				eventLogger.Errorf("Error handling event: %v", err)
				<-waitToRetry(time.Second)
				continue
			}

			eventLogger.Infof("Successfully handled event")
		}
	}
}

func isListeningEvent(evt event.Event, evtsToListen []string) bool {
	targetEvtName := evt.Name()
	for _, evtName := range evtsToListen {
		if targetEvtName == evtName {
			return true
		}
	}

	return false
}

func waitToRetry(wait time.Duration) <-chan time.Time {
	return time.After(wait)
}
