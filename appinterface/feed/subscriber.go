package feed

type Subscriber interface {
	OnNotification(n *Notification) error
}
