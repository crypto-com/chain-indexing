package feed

type Subscriber interface {
	NotifyCallback(n *Notification) error
}
