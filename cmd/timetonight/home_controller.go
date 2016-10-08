package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type homeController struct {
	App *app
}

func (c *homeController) setup(e *echo.Echo) {
	e.Get("/", c.getIndex)
}

func (c *homeController) getIndex(ctx echo.Context) error {
	durToday, err := c.App.Timer.LoggedToday()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve time logged today")
	}

	running, err := c.App.Timer.IsRunning()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve whether timer is running")
	}

	var finishTime time.Time
	durRemaining := c.App.TimePerDay - durToday
	if durRemaining > 0 {
		finishTime = time.Now().Local().Add(durRemaining)
	}

	hours := durToday / time.Hour
	minutes := (durToday - (hours * time.Hour)) / time.Minute
	data := struct {
		HoursPerDay   int
		LoggedHours   int
		LoggedMinutes int
		TimerRunning  bool
		FinishTime    time.Time
	}{
		int(c.App.TimePerDay / time.Hour),
		int(hours),
		int(minutes),
		running,
		finishTime,
	}

	return ctx.Render(http.StatusOK, "index.tmpl", data)
}
