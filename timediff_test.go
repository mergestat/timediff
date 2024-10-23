package timediff_test

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/mergestat/timediff/locale"
)

var fixtures = map[string]string{
	"-10s":                            "a few seconds ago",
	"-44s":                            "a few seconds ago",
	"-45s":                            "a minute ago",
	"-89s":                            "a minute ago",
	"-90s":                            "2 minutes ago",
	"-91s":                            "2 minutes ago",
	"-2m":                             "2 minutes ago",
	"-10m":                            "10 minutes ago",
	"-44m":                            "44 minutes ago",
	"-45m":                            "an hour ago",
	"-60m":                            "an hour ago",
	"-1h":                             "an hour ago",
	"-80m":                            "an hour ago",
	"-89m":                            "an hour ago",
	"-90m":                            "2 hours ago",
	"-2h":                             "2 hours ago",
	"-20h":                            "20 hours ago",
	"-21h":                            "21 hours ago",
	"-21h30m":                         "a day ago",
	"-22h":                            "a day ago",
	"-24h":                            "a day ago",
	"-24h30m":                         "a day ago",
	"-34h59m":                         "a day ago",
	"-35h":                            "a day ago",
	"-35h01m":                         "2 days ago",
	"-59h":                            "2 days ago",
	"-59h01m":                         "3 days ago",
	fmt.Sprintf("-%dh", 10*24):        "10 days ago",
	fmt.Sprintf("-%dh", 25*24):        "25 days ago",
	fmt.Sprintf("-%dh", 26*24):        "a month ago",
	fmt.Sprintf("-%dh", 45*24):        "a month ago",
	fmt.Sprintf("-%dh2m", 45*24):      "2 months ago",
	fmt.Sprintf("-%dh", 46*24+1):      "2 months ago",
	fmt.Sprintf("-%dh", 80*24):        "3 months ago",
	fmt.Sprintf("-%dh", 9*24*30):      "9 months ago",
	fmt.Sprintf("-%dh", 10*24*30):     "10 months ago",
	fmt.Sprintf("-%dh1m", 10*24*30):   "a year ago",
	fmt.Sprintf("-%dh", 12*24*30):     "a year ago",
	fmt.Sprintf("-%dh", 17*24*30+1):   "2 years ago",
	fmt.Sprintf("-%dh", 24*24*30):     "2 years ago",
	fmt.Sprintf("-%dh", 20*24*30*12):  "20 years ago",
	fmt.Sprintf("-%dh", 100*24*30*12): "100 years ago",

	"10s":                            "in a few seconds",
	"44s":                            "in a few seconds",
	"45s":                            "in a minute",
	"89s":                            "in a minute",
	"90s":                            "in 2 minutes",
	"2m":                             "in 2 minutes",
	"10m":                            "in 10 minutes",
	"44m":                            "in 44 minutes",
	"45m":                            "in an hour",
	"60m":                            "in an hour",
	"1h":                             "in an hour",
	"80m":                            "in an hour",
	"89m":                            "in an hour",
	"89m10s":                         "in 2 hours",
	"90m":                            "in 2 hours",
	"2h":                             "in 2 hours",
	"20h":                            "in 20 hours",
	"21h":                            "in 21 hours",
	"21h30m":                         "in a day",
	"22h":                            "in a day",
	"24h":                            "in a day",
	"35h":                            "in a day",
	"35h10m":                         "in 2 days",
	"36h":                            "in 2 days",
	"59h00m":                         "in 2 days",
	"59h01m":                         "in 3 days",
	fmt.Sprintf("%dh", 10*24):        "in 10 days",
	fmt.Sprintf("%dh", 25*24):        "in 25 days",
	fmt.Sprintf("%dh", 26*24):        "in a month",
	fmt.Sprintf("%dh", 45*24):        "in a month",
	fmt.Sprintf("%dh1m", 45*24):      "in 2 months",
	fmt.Sprintf("%dh", 46*24):        "in 2 months",
	fmt.Sprintf("%dh", 80*24):        "in 3 months",
	fmt.Sprintf("%dh", 9*24*30):      "in 9 months",
	fmt.Sprintf("%dh", 10*24*30):     "in 10 months",
	fmt.Sprintf("%dh1m", 10*24*30):   "in a year",
	fmt.Sprintf("%dh", 12*24*30):     "in a year",
	fmt.Sprintf("%dh", 24*24*30):     "in 2 years",
	fmt.Sprintf("%dh", 20*24*30*12):  "in 20 years",
	fmt.Sprintf("%dh", 100*24*30*12): "in 100 years",
}

// execFixtures is a helper function for executing tests against fixtures
func execFixtures(t *testing.T, f map[string]string, opts ...timediff.TimeDiffOption) {
	t.Helper()

	now := time.Now()

	durations := make([]string, 0, len(f))
	wants := make([]string, len(f))

	for d := range f {
		durations = append(durations, d)
	}

	sort.SliceStable(durations, func(i, j int) bool {
		pi, err := time.ParseDuration(durations[i])
		if err != nil {
			t.Fatal(err)
		}
		pj, err := time.ParseDuration(durations[j])
		if err != nil {
			t.Fatal(err)
		}
		return pi < pj
	})

	// populate the slice of formatters, corresponding to their sorted durations
	for i, d := range durations {
		wants[i] = f[d]
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)

	for d, durStr := range durations {
		want := wants[d]
		dur, err := time.ParseDuration(durStr)
		if err != nil {
			t.Fatal(err)
		}

		timeToDiff := now.Add(dur)
		got := timediff.TimeDiff(timeToDiff, opts...)

		if got != want {
			t.Fatalf("expected: %q, got: %q for duration: %q (%q)", want, got, durStr, dur)
		}

		fmt.Fprintln(w, strings.Join([]string{durStr, got}, "\t"))
	}

	if err := w.Flush(); err != nil {
		t.Fatal(err)
	}
}

func TestTimeDiff(t *testing.T) {
	execFixtures(t, fixtures, timediff.WithLocale("en-US"))
}

func TestWithStartTime(t *testing.T) {
	start := time.Date(2022, time.January, 22, 12, 0, 0, 0, time.Now().Local().Location())
	timeToDiff := time.Date(2022, time.January, 22, 10, 0, 0, 0, time.Now().Local().Location())

	want := "2 hours ago"
	got := timediff.TimeDiff(timeToDiff, timediff.WithStartTime(start))

	if got != want {
		t.Fatalf("expected: %q, got %q", want, got)
	}
}

func TestWithCustomFormatters(t *testing.T) {
	var custom = locale.Formatters{
		21 * time.Hour: func(d time.Duration) string { return fmt.Sprintf("Custom message: %.0f hours", math.Ceil(d.Hours())) },
	}

	{
		want := "Custom message: 2 hours"
		got := timediff.TimeDiff(time.Now().Add(-2*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("en-US"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}

	{
		want := "Custom message: 100 hours"
		got := timediff.TimeDiff(time.Now().Add(-100*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("en-US"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}
}
