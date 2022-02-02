package timediff

import (
	"sort"
	"time"

	"github.com/mergestat/timediff/locale"
)

// TimeDiffOption is an option used to customize a call to TimeDiff
type TimeDiffOption func(*timeDiffOptions)

type timeDiffOptions struct {
	// Start is the time to calculate the time from.
	Start time.Time

	// Locale is the locale string used by TimeDiff function.
	Locale locale.Locale

	// CustomFormatters allows the caller to override the formatters used with their own
	CustomFormatters locale.Formatters
}

// WithStartTime changes the start time from which time diff calculations are made.
// Defaults to time.Now().
func WithStartTime(t time.Time) TimeDiffOption {
	return func(opt *timeDiffOptions) {
		opt.Start = t
	}
}

// WithLocale changes the locale used for the diff operation
func WithLocale(locale locale.Locale) TimeDiffOption {
	return func(opt *timeDiffOptions) {
		opt.Locale = locale
	}
}

// WithCustomFormatters overrides the formatters supplied by any registered and used locale
func WithCustomFormatters(f locale.Formatters) TimeDiffOption {
	return func(opt *timeDiffOptions) {
		opt.CustomFormatters = f
	}
}

// TimeDiff prints a human-readable string representing the difference between two `time.Time`s.
// By default, the "start" time is time.Now(), but can be overridden with the WithStartTime(t time.Time) option.
func TimeDiff(t time.Time, options ...TimeDiffOption) string {
	// set some default options
	opt := &timeDiffOptions{Start: time.Now(), Locale: "en-US"}

	// apply the option functions
	for _, optionFn := range options {
		optionFn(opt)
	}

	var rf locale.Formatters
	if opt.CustomFormatters == nil {
		if rf = locale.Lookup(opt.Locale); rf == nil {
			return "" // TODO: find out a way to return an error in case locale is not found
		}
	} else {
		rf = opt.CustomFormatters
	}

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

	// take the diff between the supplied time and the start time
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

	return formatters[len(formatters)-1](diff)
}
