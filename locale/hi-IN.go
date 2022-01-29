package locale

import (
	"fmt"
	"math"
	"time"
)

var hindiIndia = Formatters{
	-1 << 63: func(d time.Duration) string {
		return fmt.Sprintf("%.0f वर्ष में", math.Ceil(-d.Hours()/(24.0*30*12)))
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "एक वर्ष में" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("%.0f महीने में", math.Ceil(-d.Hours()/(24.0*30)))
	},
	-45 * (24 * time.Hour): func(d time.Duration) string { return "एक महीने में" },
	-25 * (24 * time.Hour): func(d time.Duration) string {
		return fmt.Sprintf("%.0f दिन में", math.Ceil(-d.Hours()/24.0))
	},
	-35 * time.Hour:   func(_ time.Duration) string { return "एक दिन में" },
	-21 * time.Hour:   func(d time.Duration) string { return fmt.Sprintf("%.0f घंटे में", math.Ceil(-d.Hours())) },
	-89 * time.Minute: func(_ time.Duration) string { return "एक घंटे में" },
	-44 * time.Minute: func(d time.Duration) string {
		return fmt.Sprintf("%.0f मिनट में", math.Ceil(-d.Minutes()))
	},
	-89 * time.Second: func(_ time.Duration) string { return "एक मिनट में" },
	-44 * time.Second: func(_ time.Duration) string { return "कुछ ही क्षण में" },

	44 * time.Second: func(_ time.Duration) string { return "कुछ ही क्षण पहले" },
	89 * time.Second: func(_ time.Duration) string { return "एक मिनट पहले" },
	44 * time.Minute: func(d time.Duration) string {
		return fmt.Sprintf("%.0f मिनट पहले", math.Ceil(d.Minutes()))
	},
	89 * time.Minute: func(_ time.Duration) string { return "एक घंटा पहले" },
	21 * time.Hour: func(d time.Duration) string {
		return fmt.Sprintf("%.0f घंटे पहले", math.Ceil(d.Hours()))
	},
	35 * time.Hour: func(_ time.Duration) string { return "एक दिन पहले" },
	25 * (24 * time.Hour): func(d time.Duration) string {
		return fmt.Sprintf("%.0f दिन पहले", math.Ceil(d.Hours()/24.0))
	},
	45 * (24 * time.Hour): func(d time.Duration) string { return "एक महीने पहले" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("%.0f महीने पहले", math.Ceil(d.Hours()/(24.0*30)))
	},
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "एक वर्ष पहले" },
	1<<63 - 1: func(d time.Duration) string {
		return fmt.Sprintf("%.0f वर्ष पहले", math.Ceil(d.Hours()/(24.0*30*12)))
	},
}

func init() {
	Register("hi", hindiIndia) // also register it as the default hindi locale
	Register("hi-IN", hindiIndia)
}
