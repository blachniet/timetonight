package timetonight

import "time"

// Timer encapsulates interactions with a time keeping service.
type Timer interface {
	// IsRunning returns whether or not the a Toggl timer is currently running.
	IsRunning() (bool, error)
	// LoggedToday returns the amount of time that has been logged today.
	LoggedToday() (time.Duration, error)
}
