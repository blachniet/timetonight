package context

import (
	"github.com/blachniet/timetonight"
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context

	timer     timetonight.Timer
	persister timetonight.Persister
}

func (c *Context) SetTimer(t timetonight.Timer) {
	c.timer = t
}

func (c *Context) Timer() timetonight.Timer {
	return c.timer
}

func (c *Context) SetPersister(p timetonight.Persister) {
	c.persister = p
}

func (c *Context) Persister() timetonight.Persister {
	e := echo.New()
	e.Use()
	return c.persister
}

func Use(t timetonight.Timer, p timetonight.Persister) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c, t, p}
			return h(cc)
		}
	}
}
