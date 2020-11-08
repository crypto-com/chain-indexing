package feed

type Subject interface {
	Attach(s Subscriber)
	Detach(s Subscriber)
	Notify(n *Notification)
}
