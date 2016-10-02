package timetonight

import (
	"log"
	"net/http"

	"github.com/blachniet/timetonight"
)

type HandlerFunc func(t Timer, p Persister, w http.ResponseWriter, r *http.Request) (int, error)

type Handler struct {
	Timer     timetonight.Timer
	Persister timetonight.Persister
	H         HandlerFunc
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := h.H(h.Timer, h.Persister, w, r)
	if err != nil {
		log.Printf("HTTP %d: %v", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

type HandlerFactory struct {
	Timer     timetonight.Timer
	Persister timetonight.Persister
}

func (f *HandlerFactory) H(h HandlerFunc) *Handler {
	return &Handler{f.Timer, f.Persister, h}
}
