package feed

type Notification struct {
	Name    string
	Payload interface{}
}

type Feed interface {
	InitSubject() Subject
	Run() error
}
