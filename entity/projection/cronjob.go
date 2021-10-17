package projection

import "time"

// A Cron Job is a special type of projection
type CronJob interface {
	// Unique identity of the cron job.
	Id() string

	// `OnInit()` is called when the cron job first-time initializes. If an error is returned, the system will halt.
	// When you implement `onInit`, make sure all the side effects are idempotent. For example all DB changes should be
	// rollbacked on error.
	OnInit() error

	// Cron Job execution interval in duration
	Interval() time.Duration

	// Execute Cron Job logic on regular interval
	Exec() error
}
