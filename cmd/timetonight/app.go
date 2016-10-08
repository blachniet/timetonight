package main

import (
	"html/template"
	"io"
	"time"

	"github.com/blachniet/timetonight"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type app struct {
	Debug            bool
	Timer            timetonight.Timer
	Templ            *template.Template
	TemplGlobPattern string
	TimePerDay       time.Duration
}

func (a *app) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// If in debug mode, reload the templates before each render
	if a.Debug {
		t, err := template.ParseGlob(a.TemplGlobPattern)
		if err != nil {
			return errors.Wrap(err, "Err parsing templates")
		}
		a.Templ = t
	}
	return a.Templ.ExecuteTemplate(w, name, data)
}
