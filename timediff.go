package timediff

import (
	"fmt"
	"sort"
	"time"
)

// TimeDiffOption is an otion used to customize a call to TimeDiff
type TimeDiffOption func(*timeDiffOptions)

// RangeFormatters is a collection of associations between a min/max duration
// and a function for formatting an output string
type RangeFormatters map[time.Duration]func(d time.Duration) string

type timeDiffOptions struct {
	// Start is the time to calculate the time from.
	Start time.Time

	// TimeDiffRangeFormatters is the collection of duration <> formatters to use
	TimeDiffRangeFormatters RangeFormatters
}

// WithStartTime changes the start time from which time diff calculations are made.
// Defaults to time.Now().
func WithStartTime(t time.Time) TimeDiffOption {
	return func(opt *timeDiffOptions) {
		opt.Start = t
	}
}

// DefaultTimeDiffRangeFormatters are the time ranges and their corresponding string formatters.
var DefaultTimeDiffRangeFormatters = RangeFormatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("in %.0f years", -d.Hours()/(24.0*30*12)) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "in a year" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("in %.0f months", -d.Hours()/(24.0*30)) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "in a month" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("in %.0f days", -d.Hours()/24.0) },
	-35 * time.Hour:             func(_ time.Duration) string { return "in a day" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("in %.0f hours", -d.Hours()) },
	-89 * time.Minute:           func(_ time.Duration) string { return "in an hour" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("in %.0f minutes", -d.Minutes()) },
	-89 * time.Second:           func(_ time.Duration) string { return "in a minute" },
	-44 * time.Second:           func(_ time.Duration) string { return "in a few seconds" },

	44 * time.Second:           func(_ time.Duration) string { return "a few seconds ago" },
	89 * time.Second:           func(_ time.Duration) string { return "a minute ago" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f minutes ago", d.Minutes()) },
	89 * time.Minute:           func(_ time.Duration) string { return "an hour ago" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f hours ago", d.Hours()) },
	35 * time.Hour:             func(_ time.Duration) string { return "a day ago" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f days ago", d.Hours()/24.0) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "a month ago" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f months ago", d.Hours()/(24.0*30)) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "a year ago" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f years ago", d.Hours()/(24.0*30*12)) },
}

// TimeDiff prints a human-readable string representing the difference between two `time.Time`s.
// By default, the "start" time is time.Now(), but can be overridden with the WithStartTime(t time.Time) option.
func TimeDiff(t time.Time, options ...TimeDiffOption) string {
	// set some default options
	opt := &timeDiffOptions{
		Start:                   time.Now(),
		TimeDiffRangeFormatters: DefaultTimeDiffRangeFormatters,
	}

	// apply the option functions
	for _, optionFn := range options {
		optionFn(opt)
	}

	rf := opt.TimeDiffRangeFormatters

	// break up the range into slices of keys (durations) and values (formatters)
	durations := make([]time.Duration, 0, len(rf))
	formatters := make([]func(time.Duration) string, len(rf))

	// create a slice of durations
	for d := range rf {
		durations = append(durations, d)
	}

	// sort it ASC
	sort.SliceStable(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	// populate the slice of formatters, corresponding to their sorted durations
	for i, d := range durations {
		formatters[i] = rf[d]
	}

	// take the between the supplied time and the start time
	diff := opt.Start.Sub(t).Round(time.Second)

	for i := range durations {
		if diff >= 0 {
			dur := durations[i]
			if diff <= dur {
				return formatters[i](diff)
			}
		} else {
			reverseIdx := len(durations) - 1 - i
			dur := durations[reverseIdx]
			if diff >= dur {
				return formatters[reverseIdx](diff)
			}
		}
	}

	return formatters[len(formatters)](diff)
}
