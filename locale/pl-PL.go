package locale

import (
	"fmt"
	"math"
	"time"
)

var polish = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("za %.0f lat", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "za rok" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { 
		months := math.Ceil(-d.Hours()/(24.0*30))
		if months == 1 {
			return "za miesiąc"
		} else if months < 5 {
			return fmt.Sprintf("za %.0f miesiące", months)
		} else {
			return fmt.Sprintf("za %.0f miesięcy", months)
		}
	},
	-45 * (24 * time.Hour):      func(_ time.Duration) string { return "za miesiąc" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { 
		days := math.Ceil(-d.Hours()/24.0)
		if days == 1 {
			return "za dzień"
		} else {
			return fmt.Sprintf("za %.0f dni", days)
		}
	},
	-35 * time.Hour:             func(_ time.Duration) string { return "za dzień" },
	-21 * time.Hour:             func(d time.Duration) string { 
		hours := math.Ceil(-d.Hours())
		if hours == 1 {
			return "za godzinę"
		} else if hours < 5 {
			return fmt.Sprintf("za %.0f godziny", hours)
		} else {
			return fmt.Sprintf("za %.0f godzin", hours)
		}
	},
	-89 * time.Minute:           func(_ time.Duration) string { return "za godzinę" },
	-44 * time.Minute:           func(d time.Duration) string { 
		minutes := math.Ceil(-d.Minutes())
		if minutes == 1 {
			return "za minutę"
		} else if minutes < 5 {
			return fmt.Sprintf("za %.0f minuty", minutes)
		} else {
			return fmt.Sprintf("za %.0f minut", minutes)
		}
	},
	-89 * time.Second:           func(_ time.Duration) string { return "za minutę" },
	-44 * time.Second:           func(_ time.Duration) string { return "za kilka sekund" },
	44 * time.Second:           func(_ time.Duration) string { return "kilka sekund temu" },
	89 * time.Second:           func(_ time.Duration) string { return "minutę temu" },
	44 * time.Minute:           func(d time.Duration) string { 
		minutes := math.Ceil(d.Minutes())
		if minutes == 1 {
			return "minutę temu"
		} else if minutes < 5 {
			return fmt.Sprintf("%.0f minuty temu", minutes)
		} else {
			return fmt.Sprintf("%.0f minut temu", minutes)
		}
	},
	89 * time.Minute:           func(_ time.Duration) string { return "godzinę temu" },
	21 * time.Hour:             func(d time.Duration) string { 
		hours := math.Ceil(d.Hours())
		if hours == 1 {
			return "godzinę temu"
		} else if hours < 5 {
			return fmt.Sprintf("%.0f godziny temu", hours)
		} else {
			return fmt.Sprintf("%.0f godzin temu", hours)
		}
	},
	35 * time.Hour:             func(_ time.Duration) string { return "dzień temu" },
	25 * (24 * time.Hour):      func(d time.Duration) string { 
		days := math.Ceil(d.Hours()/24.0)
		if days == 1 {
			return "dzień temu"
		} else {
			return fmt.Sprintf("%.0f dni temu", days)
		}
	},
	45 * (24 * time.Hour):      func(_ time.Duration) string { return "miesiąc temu" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { 
		months := math.Ceil(d.Hours()/(24.0*30))
		if months == 1 {
			return "miesiąc temu"
		} else if months < 5 {
			return fmt.Sprintf("%.0f miesiące temu", months)
		} else {
			return fmt.Sprintf("%.0f miesięcy temu", months)
		}
	},
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "rok temu" },
	1<<63 - 1:                  func(d time.Duration) string { 
		years := math.Ceil(d.Hours()/(24.0*30*12))
		if years == 1 {
			return "rok temu"
		} else if years < 5 {
			return fmt.Sprintf("%.0f lata temu", years)
		} else {
			return fmt.Sprintf("%.0f lat temu", years)
		}
	},
}

func init() {
	Register("pl", polish)
	Register("pl-PL", polish)
}
