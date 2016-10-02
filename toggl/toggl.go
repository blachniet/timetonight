package toggl

import (
	"time"

	"github.com/blachniet/timetonight"
	"github.com/pkg/errors"
	"gopkg.in/dougEfresh/gtoggl.v8"
)

// Ensure we fully implement the Timer interface.
var _ timetonight.Timer = &Timer{}

// Timer encapsulates high-level interactions with the Toggl
// API for operations specific to this application.
type Timer struct {
	c *gtoggl.TogglClient
}

// NewTimer returns a new Toggl timer.
func NewTimer(token string) (*Timer, error) {
	c, err := gtoggl.NewClient(token)
	return &Timer{c}, errors.Wrap(err, "Err initializing Toggl API client")
}

// IsRunning returns whether a timer is currently running
// on Toggl.
func (s *Timer) IsRunning() (bool, error) {
	entries, err := s.c.TimeentryClient.List()
	if err != nil {
		return false, errors.Wrap(err, "Err retrieving time entries from Toggl")
	}

	for _, e := range entries {
		if e.Duration < 0 {
			return true, nil
		}
	}

	return false, nil
}

// LoggedToday returns the amount of time logged today in Toggl
// including the currently running timer (if applicable).
func (s *Timer) LoggedToday() (time.Duration, error) {
	panic("NotImplemented")
}
