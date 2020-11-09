package chain

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/feed"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"sync"
)

type BlockSubject struct {
	Observers sync.Map
}

func NewBlockSubject() *BlockSubject {
	return &BlockSubject{
		Observers: sync.Map{},
	}
}

func (s *BlockSubject) Attach(o feed.Subscriber) {
	s.Observers.Store(o, struct{}{})
}

func (s *BlockSubject) Notify(n *notification.BlockNotification) {
	s.Observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		if err := key.(feed.Subscriber).OnNotification(n); err != nil {
			fmt.Println("error when subscriber run callback function", err)
		}
		return true
	})
}
