package chain

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/feed"
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

func (s *BlockSubject) Detach(o feed.Subscriber) {
	s.Observers.Delete(o)
}

func (s *BlockSubject) Notify(n *feed.Notification) {
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
