package mock

import (
	"time"

	"github.com/blachniet/timetonight"
)

// Ensure we fully implement the Timer interface.
var _ timetonight.Timer = &Timer{}

type Timer struct {
	IsTimerRunningFn          func() (bool, error)
	IsTimerRunningInvoked     bool
	IsTimerRunningInvokeCount int
	LoggedTodayFn             func() (time.Duration, error)
	LoggedTodayInvoked        bool
	LoggedTodayInvokeCount    int
}

func (t *Timer) IsRunning() (bool, error) {
	t.IsTimerRunningInvoked = true
	t.IsTimerRunningInvokeCount++
	return t.IsTimerRunningFn()
}

func (s *Timer) LoggedToday() (time.Duration, error) {
	s.LoggedTodayInvoked = true
	s.LoggedTodayInvokeCount++
	return s.LoggedTodayFn()
}

func (s *Timer) Location() (*time.Location, error) {
	return time.UTC, nil
}
