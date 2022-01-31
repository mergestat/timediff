package locale

import (
	"fmt"
	"math"
	"time"
)

var germanGermany = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("in %.0f Jahren", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "in einem Jahr" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("in %.0f Monaten", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "in einem Monat" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("in %.0f Tagen", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "in einem Tag" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("in %.0f Stunden", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "in einer Stunde" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("in %.0f Minuten", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "in einer Minute" },
	-44 * time.Second:           func(_ time.Duration) string { return "in ein paar Sekunden" },

	44 * time.Second:           func(_ time.Duration) string { return "vor ein paar Sekunden" },
	89 * time.Second:           func(_ time.Duration) string { return "vor einer Minute" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("vor %.0f Minuten", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "vor einer Stunde" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("vor %.0f Stunden", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "vor einem Tag" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("vor %.0f Tagen", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "vor einem Monat" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("vor %.0f Monaten", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "vor einem Jahr" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("vor %.0f Jahren", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("de", germanGermany) // also register it as the default German locale
	Register("de-DE", germanGermany)
}
