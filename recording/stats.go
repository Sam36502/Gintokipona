package recording

import (
	"time"
)

type ActionStats struct {
	Action            *Action
	TimeSinceStart    time.Duration
	TotalFulfils      int
	ElapsedPeriods    int
	CompletionPercent float32
}

const (
	OUR_LOCALE = "Europe/Zurich"
)

func (r *RecordFile) GetActionStats(act *Action) ActionStats {
	stats := ActionStats{
		Action:       act,
		TotalFulfils: 0,
	}

	// Calculate time since start
	stats.TimeSinceStart = time.Since(act.StartDate.Time())

	// Count records
	for _, r := range r.RecordList {
		if r.ActionName == act.Name {
			stats.TotalFulfils += r.Amount
		}
	}

	// Calculate elapsed periods
	stats.ElapsedPeriods = int(stats.TimeSinceStart.Nanoseconds() / act.Period.Duration().Nanoseconds())

	// Calculate completion
	stats.CompletionPercent = (float32(stats.TotalFulfils) / float32(stats.ElapsedPeriods)) * 100

	return stats
}
