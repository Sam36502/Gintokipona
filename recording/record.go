package recording

import (
	"fmt"
	"strings"
	"time"
)

const (
	TIME_LAYOUT = "2006-01-02"
)

type RecordFile struct {
	ActionList []Action `json:"actions"`
	RecordList []Record `json:"records"`
}

type Action struct {
	Name      string       `json:"name"`
	Period    JSONDuration `json:"period"`
	StartDate JSONTime     `json:"start_date"`
}

type Record struct {
	ActionName string   `json:"action"`
	Time       JSONTime `json:"time"`
	Amount     int      `json:"amount"`
}

type JSONTime time.Time
type JSONDuration time.Duration

// Implement Marshaler and Unmarshaler interface for time wrapper
func (j *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(TIME_LAYOUT, s, time.Local)
	if err != nil {
		return err
	}
	*j = JSONTime(t)
	return nil
}

func (j JSONTime) MarshalJSON() ([]byte, error) {
	t := fmt.Sprint("\"", time.Time(j).Format(TIME_LAYOUT), "\"")
	return []byte(t), nil
}

func (j JSONTime) Time() time.Time {
	return time.Time(j)
}

// Implement Marshaler and Unmarshaler interface for duration wrapper
func (j *JSONDuration) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*j = JSONDuration(d)
	return nil
}

func (j JSONDuration) MarshalJSON() ([]byte, error) {
	d := fmt.Sprint("\"", time.Duration(j), "\"")
	return []byte(d), nil
}

func (j JSONDuration) Duration() time.Duration {
	return time.Duration(j)
}

// Gets an action from the file by its name
// Returns nil if action not found
func (r *RecordFile) GetActionByName(name string) *Action {
	for _, a := range r.ActionList {
		if a.Name == name {
			return &a
		}
	}
	return nil
}

// Gets a record's index and value from the file by when it was made
// (Only accurate to a day)
// Returns -1, nil if record not found
func (r *RecordFile) GetActionRecordByDate(act *Action, date time.Time) (int, *Record) {
	for i, e := range r.RecordList {
		if e.ActionName == act.Name && e.Time.Time().Format(TIME_LAYOUT) == date.Format(TIME_LAYOUT) {
			return i, &e
		}
	}
	return -1, nil
}
