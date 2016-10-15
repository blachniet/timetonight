package toggl

import (
	"fmt"
	"testing"
	"time"

	"gopkg.in/dougEfresh/toggl-timeentry.v8"
	"gopkg.in/dougEfresh/toggl-user.v8"
)

func TestTimerIsRunning(t *testing.T) {
	testCases := []struct {
		name           string
		durations      []int64
		expectedResult bool
		expectErr      bool
	}{
		{
			name:           "SimpleFalse",
			durations:      []int64{1234, 4567},
			expectedResult: false,
			expectErr:      false,
		},
		{
			name:           "ZeroDur",
			durations:      []int64{0, 4567},
			expectedResult: false,
			expectErr:      false,
		},
		{
			name:           "SimpleTrue",
			durations:      []int64{-1234, 4567},
			expectedResult: true,
			expectErr:      false,
		},
		{
			name:           "TrueAtEnd",
			durations:      []int64{4567, -1234},
			expectedResult: true,
			expectErr:      false,
		},
		{
			name:           "NoEntries",
			durations:      []int64{},
			expectedResult: false,
			expectErr:      false,
		},
		{
			name:           "ReturnsErr",
			durations:      []int64{},
			expectedResult: false,
			expectErr:      true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCase %v", i), func(t *testing.T) {
			mock := &mockTogglClient{
				TimeEntriesFunc: func() (gtimeentry.TimeEntries, error) {
					if tc.expectErr {
						return nil, fmt.Errorf("Fake error")
					}

					timeEntries := make([]gtimeentry.TimeEntry, len(tc.durations))
					for j, dur := range tc.durations {
						timeEntries[j] = gtimeentry.TimeEntry{Duration: dur}
					}
					return timeEntries, nil
				},
			}

			timer := Timer{mock}
			actual, err := timer.IsRunning()
			if !mock.TimeEntriesFuncInvoked {
				t.Fatalf("TimeEntries func not invoked")
			}
			if (err != nil) != (tc.expectErr) {
				t.Fatalf("Error expectation not met: expected err?: %v", tc.expectErr)
			}
			if actual != tc.expectedResult {
				t.Fatalf("Expected %v but was %v", tc.expectedResult, actual)
			}
		})
	}
}

// TODO: Still need to explicitly test server location not matching up with logged time location

func TestTimerLoggedToday(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult time.Duration
		expectErr      bool
		entries        []gtimeentry.TimeEntry
	}{
		{
			name:           "Simple",
			expectedResult: (1234 * time.Second) + (4567 * time.Second),
			expectErr:      false,
			entries: []gtimeentry.TimeEntry{
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: 1234,
				},
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: 4567,
				},
			},
		},
		{
			name:           "EntryFromYesterday",
			expectedResult: 4567 * time.Second,
			expectErr:      false,
			entries: []gtimeentry.TimeEntry{
				gtimeentry.TimeEntry{
					Start:    time.Now().Add(-24 * time.Hour),
					Duration: 1234,
				},
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: 4567,
				},
			},
		},
		{
			name:           "ZeroEntry",
			expectedResult: (4567 * time.Second),
			expectErr:      false,
			entries: []gtimeentry.TimeEntry{
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: 0,
				},
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: 4567,
				},
			},
		},
		{
			name:           "NegativeAtBeginning",
			expectedResult: (time.Duration(time.Now().Unix()-1234) * time.Second) + (4567 * time.Second),
			expectErr:      false,
			entries: []gtimeentry.TimeEntry{
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: -1234,
				},
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: 4567,
				},
			},
		},
		{
			name:           "NegativeAtEnd",
			expectedResult: (time.Duration(time.Now().Unix()-1234) * time.Second) + (4567 * time.Second),
			expectErr:      false,
			entries: []gtimeentry.TimeEntry{
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: 4567,
				},
				gtimeentry.TimeEntry{
					Start:    time.Now(),
					Duration: -1234,
				},
			},
		},
		{
			name:           "NoEntries",
			expectedResult: 0,
			expectErr:      false,
			entries:        []gtimeentry.TimeEntry{},
		},
		{
			name:           "ReturnsErr",
			expectedResult: 0,
			expectErr:      true,
			entries:        []gtimeentry.TimeEntry{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockTogglClient{
				TimeEntriesFunc: func() (gtimeentry.TimeEntries, error) {
					if tc.expectErr {
						return nil, fmt.Errorf("Fake error")
					}

					return tc.entries, nil
				},
				UserFunc: func() (*guser.User, error) {
					return &guser.User{Timezone: time.Now().Location().String()}, nil
				},
			}

			timer := Timer{mock}
			actual, err := timer.LoggedToday()
			if (err != nil) != (tc.expectErr) {
				t.Fatalf("Error expectation not met: expected err?: %v", tc.expectErr)
			}
			if actual != tc.expectedResult {
				t.Fatalf("Expected %v but was %v", tc.expectedResult, actual)
			}
		})
	}
}

func TestLocation(t *testing.T) {
	testCases := []struct {
		name      string
		loc       string
		expectErr bool
	}{
		{
			name:      "NewYork",
			loc:       "America/New_York",
			expectErr: false,
		},
		{
			name:      "Invalid",
			loc:       "abc123",
			expectErr: true,
		},
		{
			name:      "Empty",
			loc:       "",
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockTogglClient{
				UserFunc: func() (*guser.User, error) {
					return &guser.User{Timezone: tc.loc}, nil
				},
			}

			timer := Timer{mock}
			actual, err := timer.Location()
			if !mock.UserFuncInvoked {
				t.Fatalf("User func not invoked")
			}
			if (err != nil) != (tc.expectErr) {
				t.Fatalf("Error expectation not met: expected err?: %v, err: %v", tc.expectErr, err)
			}

			l, err := time.LoadLocation(tc.loc)
			if err != nil {
				if l.String() != actual.String() {
					t.Fatalf("Unexpected location: %v", *actual)
				}
			}
		})
	}
}

type mockTogglClient struct {
	UserFuncInvoked        bool
	UserFunc               func() (*guser.User, error)
	TimeEntriesFuncInvoked bool
	TimeEntriesFunc        func() (gtimeentry.TimeEntries, error)
}

func (c *mockTogglClient) User() (*guser.User, error) {
	c.UserFuncInvoked = true
	return c.UserFunc()
}

func (c *mockTogglClient) TimeEntries() (gtimeentry.TimeEntries, error) {
	c.TimeEntriesFuncInvoked = true
	return c.TimeEntriesFunc()
}
