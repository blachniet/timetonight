package timetonight

import (
	"log"
	"net/http"
)

type HandlerFunc func(t Timer, p Persister, ren Renderer, w http.ResponseWriter, r *http.Request) (int, error)

type Handler struct {
	Timer     Timer
	Persister Persister
	Renderer  Renderer
	H         HandlerFunc
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := h.H(h.Timer, h.Persister, h.Renderer, w, r)
	if err != nil {
		log.Printf("HTTP %d: %v", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

type HandlerFactory struct {
	Timer     Timer
	Persister Persister
	Renderer  Renderer
}

func (f *HandlerFactory) H(h HandlerFunc) *Handler {
	return &Handler{f.Timer, f.Persister, f.Renderer, h}
}
