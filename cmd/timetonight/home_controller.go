package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type homeController struct {
	App *app
}

func (c *homeController) setup(e *echo.Echo) {
	e.Get("/", c.getIndex)
	e.Get("/configure", c.getConfigure)
	e.Post("/configure", c.postConfigure)
}

func (c *homeController) getIndex(ctx echo.Context) error {
	timePerDay, err := c.App.Persister.TimePerDay()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve time per day")
	}

	durToday, err := c.App.Timer.LoggedToday()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve time logged today")
	}

	running, err := c.App.Timer.IsRunning()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve whether timer is running")
	}

	hours := durToday / time.Hour
	minutes := (durToday - (hours * time.Hour)) / time.Minute
	data := struct {
		HoursPerDay   int
		LoggedHours   int
		LoggedMinutes int
		TimerRunning  bool
	}{int(timePerDay / time.Hour), int(hours), int(minutes), running}

	return ctx.Render(http.StatusOK, "index.tmpl", data)
}

func (c *homeController) getConfigure(ctx echo.Context) error {
	timePerDay, err := c.App.Persister.TimePerDay()
	if err != nil {
		return errors.Wrap(err, "Err getting time per day")
	}

	togglAPIToken, err := c.App.Persister.TogglAPIToken()
	if err != nil {
		return errors.Wrap(err, "Err getting Toggl API token")
	}

	return ctx.Render(http.StatusOK, "configure.tmpl", struct {
		HoursPerDay   int
		TogglAPIToken string
	}{int(timePerDay / time.Hour), togglAPIToken})
}

func (c *homeController) postConfigure(ctx echo.Context) error {
	hrsPerDayStr := ctx.FormValue("hoursPerDay")
	hrsPerDay, err := strconv.ParseInt(hrsPerDayStr, 10, 0)
	if err != nil {
		return errors.Wrap(err, fmt.Sprint("Err converting hoursPerDay to int: ", hrsPerDayStr))
	}

	err = c.App.Persister.SetTimePerDay(time.Duration(hrsPerDay) * time.Hour)
	if err != nil {
		return errors.Wrap(err, "Err setting time per day")
	}

	err = c.App.Persister.SetTogglAPIToken(ctx.FormValue("togglAPIToken"))
	if err != nil {
		return errors.Wrap(err, "Error setting Toggl API token")
	}

	return ctx.Redirect(http.StatusFound, "/")
}
