package locale

import (
	"fmt"
	"math"
	"time"
)

var englishUnitedStates = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("in %.0f years", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "in a year" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("in %.0f months", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "in a month" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("in %.0f days", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "in a day" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("in %.0f hours", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "in an hour" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("in %.0f minutes", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "in a minute" },
	-44 * time.Second:           func(_ time.Duration) string { return "in a few seconds" },

	44 * time.Second:           func(_ time.Duration) string { return "a few seconds ago" },
	89 * time.Second:           func(_ time.Duration) string { return "a minute ago" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f minutes ago", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "an hour ago" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f hours ago", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "a day ago" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f days ago", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "a month ago" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f months ago", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "a year ago" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f years ago", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("en", englishUnitedStates) // also register it as the default english locale
	Register("en-US", englishUnitedStates)
}
