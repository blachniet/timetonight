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

	loc, err := c.App.Timer.Location()
	if err != nil {
		return errors.Wrap(err, "Error getting time zone info")
	}

	var finishTime time.Time
	remaining := c.App.TimePerDay - logged
	finishTime = time.Now().In(loc).Add(remaining)

	data := struct {
		TimerRunning bool
		TimePerDay   time.Duration
		Logged       time.Duration
		Remaining    time.Duration
		FinishTime   time.Time
	}{
		running,
		c.App.TimePerDay,
		logged,
		remaining,
		finishTime,
	}

	return ctx.Render(http.StatusOK, "index.tmpl", data)
}
