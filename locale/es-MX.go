package locale

import (
	"fmt"
	"math"
	"time"
)

var spanishMexico = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("en %.0f años", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "en un año" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("en %.0f meses", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "en un mes" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("en %.0f días", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "en un día" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("en %.0f horas", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "en una hora" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("en %.0f minutos", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "en un minuto" },
	-44 * time.Second:           func(_ time.Duration) string { return "en unos segundos" },

	44 * time.Second:           func(_ time.Duration) string { return "hace unos segundos" },
	89 * time.Second:           func(_ time.Duration) string { return "hace un minuto" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("hace %.0f minutos", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "hace una hora" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("hace %.0f horas", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "hace un día" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("hace %.0f días", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "hace un mes" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("hace %.0f meses", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "hace un año" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("hace %.0f años", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("es", spanishMexico) // also register it as the default spanish locale
	Register("es-MX", spanishMexico)
}
