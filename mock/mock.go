package mock

import (
	"time"

	"github.com/blachniet/timetonight"
)

// Ensure we fully implement the Timer interface.
var _ timetonight.Timer = &Timer{}
var _ timetonight.Persister = &Persister{}

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

type Persister struct {
	TogglAPITokenFn         func() (string, error)
	TogglAPITokenInvoked    bool
	SetTogglAPITokenFn      func(string) error
	SetTogglAPITokenInvoked bool
	TimePerDayFn            func() (time.Duration, error)
	TimePerDayInvoked       bool
	SetTimePerDayFn         func(time.Duration) error
	SetTimePerDayInvoked    bool
}

func (p *Persister) TogglAPIToken() (string, error) {
	p.TogglAPITokenInvoked = true
	return p.TogglAPITokenFn()
}

func (p *Persister) SetTogglAPIToken(s string) error {
	p.SetTogglAPITokenInvoked = true
	return p.SetTogglAPITokenFn(s)
}

func (p *Persister) TimePerDay() (time.Duration, error) {
	p.TimePerDayInvoked = true
	return p.TimePerDayFn()
}

func (p *Persister) SetTimePerDay(t time.Duration) error {
	p.SetTimePerDayInvoked = true
	return p.SetTimePerDayFn(t)
}
