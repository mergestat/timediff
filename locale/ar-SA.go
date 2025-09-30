package locale

import (
	"fmt"
	"math"
	"time"
)

func arabicCount(value float64, singular, dual, few, many string) string {
	n := int(math.Ceil(value))
	if n < 0 {
		n = -n
	}

	switch {
	case n == 1:
		return singular
	case n == 2:
		return dual
	case n%100 >= 3 && n%100 <= 10:
		return fmt.Sprintf("%d %s", n, few)
	default:
		return fmt.Sprintf("%d %s", n, many)
	}
}

func arabicSince(value float64, singular, dual, few, many string) string {
	return fmt.Sprintf("منذ %s", arabicCount(value, singular, dual, few, many))
}

func arabicIn(value float64, singular, dual, few, many string) string {
	return fmt.Sprintf("بعد %s", arabicCount(value, singular, dual, few, many))
}

var arabicSaudi = Formatters{
	-1 << 63: func(d time.Duration) string {
		return arabicIn(-d.Hours()/(24.0*30*12), "سنة", "سنتين", "سنوات", "سنة")
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "بعد سنة" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return arabicIn(-d.Hours()/(24.0*30), "شهر", "شهرين", "أشهر", "شهر")
	},
	-45 * (24 * time.Hour): func(_ time.Duration) string { return "بعد شهر" },
	-25 * (24 * time.Hour): func(d time.Duration) string {
		return arabicIn(-d.Hours()/24.0, "يوم", "يومين", "أيام", "يوم")
	},
	-35 * time.Hour: func(_ time.Duration) string { return "بعد يوم" },
	-21 * time.Hour: func(d time.Duration) string {
		return arabicIn(-d.Hours(), "ساعة", "ساعتين", "ساعات", "ساعة")
	},
	-89 * time.Minute: func(_ time.Duration) string { return "بعد ساعة" },
	-44 * time.Minute: func(d time.Duration) string {
		return arabicIn(-d.Minutes(), "دقيقة", "دقيقتين", "دقائق", "دقيقة")
	},
	-89 * time.Second: func(_ time.Duration) string { return "بعد دقيقة" },
	-44 * time.Second: func(_ time.Duration) string { return "بعد ثوانٍ قليلة" },

	44 * time.Second: func(_ time.Duration) string { return "منذ ثوانٍ قليلة" },
	89 * time.Second: func(_ time.Duration) string { return "منذ دقيقة" },
	44 * time.Minute: func(d time.Duration) string {
		return arabicSince(d.Minutes(), "دقيقة", "دقيقتين", "دقائق", "دقيقة")
	},
	89 * time.Minute: func(_ time.Duration) string { return "منذ ساعة" },
	21 * time.Hour: func(d time.Duration) string {
		return arabicSince(d.Hours(), "ساعة", "ساعتين", "ساعات", "ساعة")
	},
	35 * time.Hour: func(_ time.Duration) string { return "منذ يوم" },
	25 * (24 * time.Hour): func(d time.Duration) string {
		return arabicSince(d.Hours()/24.0, "يوم", "يومين", "أيام", "يوم")
	},
	45 * (24 * time.Hour): func(_ time.Duration) string { return "منذ شهر" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return arabicSince(d.Hours()/(24.0*30), "شهر", "شهرين", "أشهر", "شهر")
	},
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "منذ سنة" },
	1<<63 - 1: func(d time.Duration) string {
		return arabicSince(d.Hours()/(24.0*30*12), "سنة", "سنتين", "سنوات", "سنة")
	},
}

func init() {
	Register("ar", arabicSaudi)
	Register("ar-SA", arabicSaudi)
}
