package feed

import "github.com/crypto-com/chainindex/infrastructure/notification"

type Subscriber interface {
	OnNotification(n *notification.BlockNotification) error
}
