package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"

	"github.com/blachniet/timetonight"
	"github.com/blachniet/timetonight/bolt"
	"github.com/blachniet/timetonight/toggl"
)

type app struct {
	Debug     bool
	Timer     timetonight.Timer
	Persister timetonight.Persister
}

func (a *app) trySetTimer() (bool, error) {
	togglAPIToken, err := a.Persister.TogglAPIToken()
	if err != nil {
		return false, errors.Wrap(err, "Err retrieving Toggl API token")
	}

	if togglAPIToken == "" {
		return false, nil
	}

	timer, err := toggl.NewTimer(togglAPIToken)
	if err != nil {
		return false, errors.Wrap(err, "Err connecting Toggl timer")
	}

	a.Timer = timer
	return true, nil
}

func main() {
	debug := flag.Bool("debug", false, "Enables debugging")
	tmplPattern := flag.String("templates", "./templates/*.tmpl", "Glob pattern for templates")
	dbPath := flag.String("db", "/Users/brian.lachniet/timetonight.db", "Path to bolt database file")
	flag.Parse()

	persister := bolt.NewPersister(*dbPath)
	err := persister.Open()
	if err != nil {
		log.Fatalf("Err opening persister: %+v", err)
	}
	defer persister.Close()

	app := &app{
		Debug:     debug != nil && *debug,
		Timer:     nil,
		Persister: persister,
	}

	// Echo Setup
	e := echo.New()
	e.SetDebug(app.Debug)
	e.SetRenderer(&renderer{
		debug:   app.Debug,
		pattern: *tmplPattern,
	})

	// Echo Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		// Redirect to configuration page if not configured
		return func(c echo.Context) error {
			if c.Path() != "/configure" && app.Timer == nil {
				ok, err := app.trySetTimer()
				if err != nil {
					errors.Wrap(err, "Error setting timer")
				}
				if !ok {
					return c.Redirect(http.StatusFound, "/configure")
				}
			}
			return h(c)
		}
	})

	// Controllers
	homeController := &homeController{app}
	homeController.setup(e)

	// Echo Run
	e.Run(standard.New(":3000"))
}
