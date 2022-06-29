package recording

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type RecordFile struct {
	ActionList []Action `json:"actions"`
	RecordList []Record `json:"records"`
}

type Action struct {
	Name         string       `json:"name"`
	Participants []string     `json:"participants"`
	Period       JSONDuration `json:"period"`
	StartDate    JSONTime     `json:"start_date"`
}

type Record struct {
	ActionName string   `json:"action"`
	Time       JSONTime `json:"time"`
}

type JSONTime time.Time
type JSONDuration time.Duration

// Implement Marshaler and Unmarshaler interface for time wrapper
func (j *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JSONTime(t)
	return nil
}

func (j JSONTime) MarshalJSON() ([]byte, error) {
	t := fmt.Sprint(time.Time(j))
	return []byte(t), nil
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
	d := fmt.Sprint(time.Duration(j))
	return []byte(d), nil
}

func (r *RecordFile) GetActionNames() []string {
	actions := []string{}
	for _, a := range r.ActionList {
		actions = append(actions, a.Name)
	}
	sort.Strings(actions)
	return actions
}

func (r *RecordFile) GetParticipants() []string {
	partmap := map[string]bool{}
	for _, a := range r.ActionList {
		for _, p := range a.Participants {
			partmap[p] = true
		}
	}
	participants := make([]string, 0)
	for p := range partmap {
		participants = append(participants, p)
	}

	sort.Strings(participants)
	return participants
}
