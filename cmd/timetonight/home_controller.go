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
	logged, err := c.App.Timer.LoggedToday()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve time logged today")
	}

	running, err := c.App.Timer.IsRunning()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve whether timer is running")
	}

	var finishTime time.Time
	remaining := c.App.TimePerDay - logged
	if remaining > 0 {
		finishTime = time.Now().Local().Add(remaining)
	}

	data := struct {
		TimerRunning  bool
		TimePerDay    time.Duration
		LoggedTime    time.Duration
		RemainingTime time.Duration
		FinishTime    time.Time
	}{
		running,
		c.App.TimePerDay,
		logged,
		remaining,
		finishTime,
	}

	return ctx.Render(http.StatusOK, "index.tmpl", data)
}
