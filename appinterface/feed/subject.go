package feed

type Notification struct {
	Name    string
	Payload interface{}
}

type Subject interface {
	Attach(s Subscriber)
	Detach(s Subscriber)
	Notify(n *Notification)
}
