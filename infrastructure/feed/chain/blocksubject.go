package chain

import (
	"sync"

	applogger "github.com/crypto-com/chainindex/internal/logger"

	event_interface "github.com/crypto-com/chainindex/appinterface/event"
	"github.com/crypto-com/chainindex/infrastructure/notification"
)

type BlockSubject struct {
	logger applogger.Logger

	Observers sync.Map
}

func NewBlockSubject(logger applogger.Logger) *BlockSubject {
	return &BlockSubject{
		logger: logger.WithFields(applogger.LogFields{
			"module": "BlockSubject",
		}),

		Observers: sync.Map{},
	}
}

func (s *BlockSubject) Attach(subj *BlockSubscriber) {
	s.Observers.Store(subj, struct{}{})
}

func (s *BlockSubject) Notify(n *notification.BlockNotification, eventStore *event_interface.RDbStore) {
	s.Observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		if err := key.(*BlockSubscriber).OnNotification(n, s.logger, eventStore); err != nil {
			s.logger.Errorf("error when subscriber run callback function: %v", err)
		}
		return true
	})
}
