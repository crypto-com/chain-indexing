package eventhandler

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
)

var _ Handler = &ProjectionHandler{}

type ProjectionHandler struct {
	logger     applogger.Logger
	projection projection_entity.Projection
}

func NewProjectionHandler(logger applogger.Logger, projection projection_entity.Projection) *ProjectionHandler {
	return &ProjectionHandler{
		logger,
		projection,
	}
}

func (handler *ProjectionHandler) GetLastHandledEventHeight() (*int64, error) {
	return handler.projection.GetLastHandledEventHeight()
}

func (handler *ProjectionHandler) HandleEvents(blockHeight int64, events []event.Event) error {
	logger := handler.logger.WithFields(applogger.LogFields{
		"height": blockHeight,
	})

	filteredEvents := make([]event.Event, 0)
	for _, event := range events {
		if !isListeningEvent(event, handler.projection.GetEventsToListen()) {
			//eventLogger.WithFields(applogger.LogFields{
			//	"eventName": event.Name(),
			//}).Debugf("skipping because event is not one of the listening events")
			continue
		}
		filteredEvents = append(filteredEvents, event)
	}

	logger = logger.WithFields(applogger.LogFields{
		"eventCount": len(filteredEvents),
	})
	if err := handler.projection.HandleEvents(blockHeight, filteredEvents); err != nil {
		logger.WithFields(applogger.LogFields{
			"events": filteredEvents,
		}).Errorf("error handling filtered events: %v", err)
		return fmt.Errorf("error handling filtered events: %v", err)
	}

	logger.Infof("successfully handled events")
	return nil
}

func (handler *ProjectionHandler) Id() string {
	return handler.projection.Id()
}

func isListeningEvent(event event.Event, eventsToListen []string) bool {
	targetEventName := event.Name()
	for _, eventName := range eventsToListen {
		if targetEventName == eventName {
			return true
		}
	}

	return false
}
