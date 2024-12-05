package locale

import (
	"fmt"
	"math"
	"time"
)

var swedishSweden = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("om %.0f år", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "om ett år" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("om %.0f månader", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "om en månad" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("om %.0f dagar", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "om en dag" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("om %.0f timmar", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "om en timme" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("om %.0f minuter", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "om en minut" },
	-44 * time.Second:           func(_ time.Duration) string { return "om några sekunder" },

	44 * time.Second:           func(_ time.Duration) string { return "några sekunder sedan" },
	89 * time.Second:           func(_ time.Duration) string { return "en minut sedan" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f minuter sedan", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "en timme sedan" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f timmar sedan", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "en dag sedan" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f dagar sedan", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "en månad sedan" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f månader sedan", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "ett år sedan" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f år sedan", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("sv", swedishSweden) // also register it as the default swedish locale
	Register("sv-SE", swedishSweden)
}
