package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blachniet/timetonight"
	"github.com/blachniet/timetonight/mock"
	"gopkg.in/dougEfresh/gtoggl.v8"
)

func main() {
	targetDuration := 9 * time.Hour
	togglToken := os.Getenv("TIMETON_TOGGLTOKEN")
	if togglToken == "" {
		log.Fatalf("No Toggl API token found in env variable TIMETON_TOGGLTOKEN")
	}

	tc, err := gtoggl.NewClient(togglToken)
	if err != nil {
		log.Fatalf("Failed to connect to Toggl: %v", err)
	}

	entries, err := tc.TimeentryClient.List()
	if err != nil {
		log.Fatalf("Err getting Toggl time entries: %v", err)
	}

	var todayDur time.Duration
	var timerRunning bool
	for _, entry := range entries {
		nowYear, nowMonth, nowDay := time.Now().Local().Date()
		startYear, startMonth, startDay := entry.Start.Local().Date()
		if nowYear == startYear && nowMonth == startMonth && nowDay == startDay {
			if entry.Duration > 0 {
				todayDur += time.Duration(entry.Duration) * time.Second
			} else {
				timerRunning = true
				todayDur += (time.Duration(time.Now().UTC().Unix()) * time.Second) + (time.Duration(entry.Duration) * time.Second)
			}
		}
	}

	fmt.Print("Timer    :")
	if timerRunning {
		fmt.Println(" On")
	} else {
		fmt.Println(" Off")
	}

	hours := todayDur / time.Hour
	minutes := (todayDur - (hours * time.Hour)) / time.Minute
	fmt.Printf("Logged   : %dh %dm\n", hours, minutes)

	remainingDur := targetDuration - todayDur
	if remainingDur > 0 {
		remHours := remainingDur / time.Hour
		remMin := (remainingDur - (remHours * time.Hour)) / time.Minute
		fmt.Printf("Remaining: %dh %dm\n", remHours, remMin)

		if timerRunning {
			fmt.Printf("=== You should be done around %v\n", time.Now().Local().Add(remainingDur))
		} else {
			fmt.Printf("=== If you start your timer now you should be done around %v\n", time.Now().Local().Add(remainingDur))
		}
	} else {
		fmt.Println("=== You're done!")
	}

	timer := &mock.Timer{}
	persister := &mock.Persister{}
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
	return http.StatusOK, ren.Render(w, "index.tmpl", nil)
}
