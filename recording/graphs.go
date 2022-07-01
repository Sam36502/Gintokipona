package recording

import (
	"fmt"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	BARCHART_FILENAME = "barchart.html"
)

func CreateBarChart(rec *RecordFile, act *Action) {
	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Fulfilments per Day",
			Subtitle: "That is, how many times you fulfilled your action per day",
		}),
	)

	f, err := os.Create(BARCHART_FILENAME)
	if err != nil {
		fmt.Printf("[Error] Failed to create bar-chart '%s':\n%s\n", BARCHART_FILENAME, err.Error())
		return
	}
	defer f.Close()

	// Create X Axis
	xTitles := []string{}
	records := []opts.BarData{}
	for d := act.StartDate.Time(); d.Before(time.Now()); d = d.Add(24 * time.Hour) {
		xTitles = append(xTitles, d.Format(TIME_LAYOUT))
		_, r := rec.GetActionRecordByDate(act, d)
		if r == nil {
			records = append(records, opts.BarData{Value: 0})
		} else {
			records = append(records, opts.BarData{Value: r.Amount})
		}
	}
	bar.SetXAxis(xTitles).AddSeries(
		fmt.Sprintf("Action '%s' Fulfilments", act.Name),
		records,
	)
	bar.Render(f)
}
