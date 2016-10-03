package timetonight

import (
	"html/template"
	"io"

	"github.com/pkg/errors"
)

type Renderer interface {
	Render(wr io.Writer, name string, data interface{}) error
}

type DefaultRender struct {
	t       *template.Template
	pattern string
	debug   bool
}

func NewDefaultRenderer(pattern string, debug bool) (*DefaultRender, error) {
	t, err := template.ParseGlob(pattern)
	if err != nil {
		return nil, errors.Wrap(err, "Err parsing templates")
	}

	return &DefaultRender{t, pattern, debug}, nil
}

func (r *DefaultRender) Render(wr io.Writer, name string, data interface{}) error {
	// If in debug mode, reload the templates before each render
	if r.debug {
		t, err := template.ParseGlob(r.pattern)
		if err != nil {
			return errors.Wrap(err, "Err parsing templates")
		}
		r.t = t
	}
	return r.t.ExecuteTemplate(wr, name, data)
}
