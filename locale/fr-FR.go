package locale

import (
	"fmt"
	"math"
	"time"
)

var frenchFrance = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("dans %.0f ans", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "dans un an" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("dans %.0f mois", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "dans un mois" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("dans %.0f jours", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "dans un jour" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("dans %.0f heures", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "dans une heure" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("dans %.0f minutes", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "dans une minute" },
	-44 * time.Second:           func(_ time.Duration) string { return "dans quelques secondes" },

	44 * time.Second:           func(_ time.Duration) string { return "il y a quelques secondes" },
	89 * time.Second:           func(_ time.Duration) string { return "il y a une minute" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("il y a %.0f minutes", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "il y a une heure" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("il y a %.0f heures", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "il y a un jour" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("il y a %.0f jours", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "il y a un mois" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("il y a %.0f mois", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "il y a un an" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("il y a %.0f ans", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("fr", frenchFrance)
	Register("fr-FR", frenchFrance)
}
