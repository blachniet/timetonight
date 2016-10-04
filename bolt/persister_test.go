package bolt

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func getTestPersister(t *testing.T) (*os.File, *Persister) {
	tmpfile, err := ioutil.TempFile("", "testdb")
	if err != nil {
		t.Fatal("Failed to create temp database ", err)
	}
	tmpfile.Close()

	p := NewPersister(tmpfile.Name())
	err = p.Open()
	if err != nil {
		t.Fatal("Failed to open temp database ", err)
	}
	return tmpfile, p
}

func TestTogglAPIToken(t *testing.T) {
	tmpfile, p := getTestPersister(t)
	defer os.Remove(tmpfile.Name())
	defer p.Close()

	token, err := p.TogglAPIToken()
	if err != nil {
		t.Fatalf("Failed to retrieve token")
	}
	if token != "" {
		t.Fatalf("Expected empty token but was %v", token)
	}

	err = p.SetTogglAPIToken("abc123")
	if err != nil {
		t.Fatalf("Error setting toggl api token")
	}

	token, err = p.TogglAPIToken()
	if err != nil {
		t.Fatalf("Failed to retrieve token")
	}
	if token != "abc123" {
		t.Fatalf("Expected 'abc123' token but was %v", token)
	}

	err = p.SetTogglAPIToken("")
	if err != nil {
		t.Fatalf("Error setting toggl api token to empty value")
	}

	token, err = p.TogglAPIToken()
	if err != nil {
		t.Fatalf("Failed to retrieve token")
	}
	if token != "" {
		t.Fatalf("Expected empty token again but was %v", token)
	}
}

func TestTimePerDay(t *testing.T) {
	tmpfile, p := getTestPersister(t)
	defer os.Remove(tmpfile.Name())
	defer p.Close()

	tpd, err := p.TimePerDay()
	if err != nil {
		t.Fatalf("Failed to retrieve time per day")
	}
	if tpd != 0 {
		t.Fatalf("Expected empty time per day but was %v", tpd)
	}

	err = p.SetTimePerDay(time.Hour * 9)
	if err != nil {
		t.Fatalf("Error setting time per day api token")
	}

	tpd, err = p.TimePerDay()
	if err != nil {
		t.Fatalf("Failed to retrieve time per day")
	}
	if tpd != time.Hour*9 {
		t.Fatalf("Expected 'abc123' time per day but was %v", tpd)
	}

	err = p.SetTimePerDay(0)
	if err != nil {
		t.Fatalf("Error setting time per day api token to empty value")
	}

	tpd, err = p.TimePerDay()
	if err != nil {
		t.Fatalf("Failed to retrieve time per day")
	}
	if tpd != 0 {
		t.Fatalf("Expected empty time per day again but was %v", tpd)
	}
}
