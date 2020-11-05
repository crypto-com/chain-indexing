package projection

import (
	"github.com/crypto-com/chainindex/ddd/event"
)

// Projection interface defines the necessary methods of to create a projection
type Projection interface {
	// Unique identity of the projection.
	Id() string

	// Returns an array of event names to listen. All versions of the events will be listened.
	GetEventsToListen() []string

	// Returns the last handled event sequence. nil mean no event has been handled
	GetLastHandledEventSeq() *int64

	// `OnInit()` is called when the projection first-time initializes (Before the first event is
	// handled). If an error is returned, the system will attempt to run again on next restart and
	// no event will be handle by this projection until this method return no error.
	// When you implement `onInit`, make sure all the side effects are idempotent. For example all
	// DB changes should be rollbacked on error.
	OnInit() error

	// Handle events that matches `GetEventsToListen()` and create projection
	HandleEvent(event event.Event) error
}
