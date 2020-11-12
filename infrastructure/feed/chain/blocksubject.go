package chain

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/rdb"
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

func (s *BlockSubject) Attach(subj *BlockSubscriber) {
	s.Observers.Store(subj, struct{}{})
}

func (s *BlockSubject) Notify(n *notification.BlockNotification, handle *rdb.Handle) {
	s.Observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		if err := key.(*BlockSubscriber).OnNotification(n, handle); err != nil {
			fmt.Println("error when subscriber run callback function", err)
		}
		return true
	})
}
