package feed

import "github.com/crypto-com/chainindex/infrastructure/notification"

type Subject interface {
	Attach(s Subscriber)
	Notify(n *notification.BlockNotification)
}
