package locale

import (
	"fmt"
	"math"
	"time"
)

var turkishTurkey = Formatters{
	-1 << 63: func(d time.Duration) string {
		return fmt.Sprintf("%.0f yıl içinde", math.Ceil(-d.Hours()/(24.0*30*12)))
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "bir yıl içinde" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f ay içinde", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "bir ay içinde" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f gün içinde", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "bir gün içinde" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f saat içinde", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "bir saat içinde" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f dakika içinde", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "bir dakika içinde" },
	-44 * time.Second:           func(_ time.Duration) string { return "birkaç saniye içinde" },

	44 * time.Second:           func(_ time.Duration) string { return "birkaç saniye önce" },
	89 * time.Second:           func(_ time.Duration) string { return "bir dakika önce" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f dakika önce", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "bir saat önce" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f saat önce", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "bir gün önce" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f gün önce", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "bir ay önce" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f ay önce", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "bir yıl önce" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f yıl önce", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("tr", turkishTurkey) // also register it as the default turkish locale
	Register("tr-TR", turkishTurkey)
}
