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
func (t *Timer) IsRunning() (bool, error) {
	entries, err := t.c.TimeentryClient.List()
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
func (t *Timer) LoggedToday() (time.Duration, error) {
	var dur time.Duration
	entries, err := t.c.TimeentryClient.List()
	if err != nil {
		return dur, errors.Wrap(err, "Err retrieving time entries from Toggl")
	}

	loc, err := t.Location()
	if err != nil {
		return dur, errors.Wrap(err, "Err getting location for Toggl user")
	}

	nowYear, nowMonth, nowDay := time.Now().In(loc).Date()
	for _, e := range entries {
		startYear, startMonth, startDay := e.Start.In(loc).Date()
		if nowYear == startYear && nowMonth == startMonth && nowDay == startDay {
			if e.Duration > 0 {
				dur += time.Duration(e.Duration) * time.Second
			} else {
				dur += (time.Duration(time.Now().UTC().Unix()) * time.Second) + (time.Duration(e.Duration) * time.Second)
			}
		}
	}

	return dur, nil
}

// Location returns time zone information for the user associated
// with this timer.
func (t *Timer) Location() (*time.Location, error) {
	u, err := t.c.UserClient.Get(false)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting Toggl user")
	}

	loc, err := time.LoadLocation(u.Timezone)
	return loc, errors.Wrap(err, "Error parsing timezone")
}
