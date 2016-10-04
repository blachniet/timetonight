package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blachniet/timetonight"
	"github.com/blachniet/timetonight/mock"
	"github.com/blachniet/timetonight/toggl"
)

func main() {
	togglToken := os.Getenv("TOGGL_API_TOKEN")

	persister := &mock.Persister{}
	persister.TogglAPITokenFn = func() (string, error) {
		return togglToken, nil
	}
	persister.TimePerDayFn = func() (time.Duration, error) {
		return time.Hour * 9, nil
	}

	timer, err := toggl.NewTimer(togglToken)
	if err != nil {
		log.Fatalf("Err creating Toggl timer: %+v", err)
	}

	ren, err := timetonight.NewDefaultRenderer("./templates/*.tmpl", true)
	if err != nil {
		log.Fatalf("Err creating renderer: %+v", err)
	}

	hf := &timetonight.HandlerFactory{
		Timer:     timer,
		Persister: persister,
		Renderer:  ren,
	}

	http.Handle("/", hf.H(getIndex))
	http.ListenAndServe(":3000", nil)
}

func getIndex(t timetonight.Timer, p timetonight.Persister, ren timetonight.Renderer,
	w http.ResponseWriter, r *http.Request) (int, error) {

	timePerDay, err := p.TimePerDay()
	if err != nil {
		log.Printf("Failed to retrieve time per day: %+v", err)
		return http.StatusInternalServerError, err
	}

	durToday, err := t.LoggedToday()
	if err != nil {
		log.Printf("Failed to retrieve time logged today: %+v", err)
		return http.StatusInternalServerError, err
	}

	running, err := t.IsRunning()
	if err != nil {
		log.Printf("Failed to retrieve whether timer is running: %+v", err)
		return http.StatusInternalServerError, err
	}

	hours := durToday / time.Hour
	minutes := (durToday - (hours * time.Hour)) / time.Minute
	data := struct {
		HoursPerDay   int
		LoggedHours   int
		LoggedMinutes int
		TimerRunning  bool
	}{int(timePerDay / time.Hour), int(hours), int(minutes), running}

	return http.StatusOK, ren.Render(w, "index.tmpl", data)
}
