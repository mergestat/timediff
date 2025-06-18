package locale

import (
	"fmt"
	"math"
	"time"
)

var farsiIran = Formatters{
	-1 << 63: func(d time.Duration) string {
		return fmt.Sprintf("در %.0f سال", math.Ceil(-d.Hours()/(24.0*30*12)))
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "در یک سال" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("در %.0f ماه", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "در یک ماه" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("در %.0f روز", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "در یک روز" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("در %.0f ساعت", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "در یک ساعت" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("در %.0f دقیقه", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "در یک دقیقه" },
	-44 * time.Second:           func(_ time.Duration) string { return "در چند ثانیه" },

	44 * time.Second:           func(_ time.Duration) string { return "چند ثانیه پیش" },
	89 * time.Second:           func(_ time.Duration) string { return "یک دقیقه پیش" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f دقیقه پیش", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "یک ساعت پیش" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f ساعت پیش", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "یک روز پیش" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f روز پیش", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "یک ماه پیش" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f ماه پیش", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "یک سال پیش" },
	1<<63 - 1: func(d time.Duration) string {
		return fmt.Sprintf("%.0f سال پیش", math.Ceil(d.Hours()/(24.0*30*12)))
	},
}

func init() {
	Register("fa", farsiIran) // also register it as the default english locale
	Register("fa-IR", farsiIran)
}
