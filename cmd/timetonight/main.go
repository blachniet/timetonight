package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"

	"github.com/blachniet/timetonight/bolt"
	"github.com/blachniet/timetonight/context"
	"github.com/blachniet/timetonight/toggl"
)

func main() {
	persister := bolt.NewPersister("/Users/brian.lachniet/timetonight.db")
	err := persister.Open()
	if err != nil {
		log.Fatalf("Err opening persister: %+v", err)
	}
	defer persister.Close()

	// Echo Setup
	e := echo.New()
	e.SetDebug(true)
	e.SetRenderer(&renderer{
		debug:   e.Debug(),
		pattern: "./templates/*.tmpl",
	})

	// Echo Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(context.Use(nil, persister))
	e.Use(checkConfigurationRequired)

	// Echo Handlers
	e.Get("/", getIndex)
	e.Get("/configure", getConfigure)
	e.Post("/configure", postConfigure)

	// Echo Run
	e.Run(standard.New(":3000"))
}

func checkConfigurationRequired(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*context.Context)
		if c.Path() != "/configure" {
			if cc.Timer() == nil {
				return c.Redirect(http.StatusFound, "/configure")
			}
		}
		return h(c)
	}
}

func getConfigure(c echo.Context) error {
	cc := c.(*context.Context)
	timePerDay, err := cc.Persister().TimePerDay()
	if err != nil {
		return errors.Wrap(err, "Err getting time per day")
	}

	togglAPIToken, err := cc.Persister().TogglAPIToken()
	if err != nil {
		return errors.Wrap(err, "Err getting Toggl API token")
	}

	return c.Render(http.StatusOK, "configure.tmpl", struct {
		HoursPerDay   int
		TogglAPIToken string
	}{int(timePerDay / time.Hour), togglAPIToken})
}

func postConfigure(c echo.Context) error {
	cc := c.(*context.Context)
	log.Println("Post configure")
	togglAPIToken := c.FormValue("togglAPIToken")
	if togglAPIToken != "" {
		log.Println("Setting token ", togglAPIToken)
		t, err := toggl.NewTimer(togglAPIToken)
		if err != nil {
			return errors.Wrap(err, "Err creating toggl timer")
		}
		cc.SetTimer(t)
	} else {
		cc.SetTimer(nil)
	}

	return c.Redirect(http.StatusFound, "/")
}

func getIndex(c echo.Context) error {
	cc := c.(*context.Context)
	timePerDay, err := cc.Persister().TimePerDay()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve time per day")
	}

	durToday, err := cc.Timer().LoggedToday()
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve time logged today")
	}

	running, err := cc.Timer().IsRunning()
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

	return c.Render(http.StatusOK, "index.tmpl", data)
}
