package locale

import (
	"fmt"
	"math"
	"time"
)

var indonesianIndonesia = Formatters{
	-1 << 63: func(d time.Duration) string {
		return fmt.Sprintf("dalam %.0f tahun", math.Ceil(-d.Hours()/(24.0*30*12)))
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "dalam setahun" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("dalam %.0f bulan", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "dalam sebulan" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("dalam %.0f hari", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "dalam sehari" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("dalam %.0f jam", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "dalam satu jam" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("dalam %.0f menit", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "dalam satu menit" },
	-44 * time.Second:           func(_ time.Duration) string { return "dalam beberapa detik" },

	44 * time.Second:      func(_ time.Duration) string { return "beberapa detik yang lalu" },
	89 * time.Second:      func(_ time.Duration) string { return "satu menit yang lalu" },
	44 * time.Minute:      func(d time.Duration) string { return fmt.Sprintf("%.0f menit yang lalu", math.Ceil(d.Minutes())) },
	89 * time.Minute:      func(_ time.Duration) string { return "satu jam yang lalu" },
	21 * time.Hour:        func(d time.Duration) string { return fmt.Sprintf("%.0f jam yang lalu", math.Ceil(d.Hours())) },
	35 * time.Hour:        func(_ time.Duration) string { return "sehari yang lalu" },
	25 * (24 * time.Hour): func(d time.Duration) string { return fmt.Sprintf("%.0f hari yang lalu", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour): func(d time.Duration) string { return "sebulan yang lalu" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("%.0f bulan yang lalu", math.Ceil(d.Hours()/(24.0*30)))
	},
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "setahun yang lalu" },
	1<<63 - 1: func(d time.Duration) string {
		return fmt.Sprintf("%.0f tahun yang lalu", math.Ceil(d.Hours()/(24.0*30*12)))
	},
}

func init() {
	Register("id", indonesianIndonesia) // juga mendaftarkan sebagai lokal bahasa Indonesia default
	Register("id-ID", indonesianIndonesia)
}
