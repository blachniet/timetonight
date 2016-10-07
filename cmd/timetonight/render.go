package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type renderer struct {
	t       *template.Template
	pattern string
	debug   bool
}

func (r *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// If in debug mode, reload the templates before each render
	if r.debug {
		t, err := template.ParseGlob(r.pattern)
		if err != nil {
			return errors.Wrap(err, "Err parsing templates")
		}
		r.t = t
	}
	return r.t.ExecuteTemplate(w, name, data)
}
