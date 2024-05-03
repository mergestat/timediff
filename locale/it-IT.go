package locale

import (
	"fmt"
	"math"
	"time"
)

var italianItaly = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("tra %.0f anni", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "tra un anno" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("tra %.0f mesi", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "tra un mese" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("tra %.0f giorni", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "domani" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("tra %.0f ore", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "tra un'ora" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("tra %.0f minuti", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "tra un minuto" },
	-44 * time.Second:           func(_ time.Duration) string { return "tra pochi secondi" },

	44 * time.Second:           func(_ time.Duration) string { return "pochi secondi fa" },
	89 * time.Second:           func(_ time.Duration) string { return "un minuto fa" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f minuti fa", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "un'ora fa" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f ore fa", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "ieri" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f giorni fa", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "un mese fa" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f mesi fa", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "un anno fa" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f anni fa", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("it", italianItaly)
	Register("it-IT", italianItaly)
}
