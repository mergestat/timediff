package locale

import (
	"fmt"
	"math"
	"time"
)

func arabicPlural(n int, singular, dual, few, many string) string {
	switch n {
	case 1:
		return singular
	case 2:
		return dual
	default:
		if n <= 10 {
			return fmt.Sprintf(few, n)
		}
		return fmt.Sprintf(many, n)
	}
}

func arabicPastMinutes(n int) string {
	return arabicPlural(n, "منذ دقيقة واحدة", "منذ دقيقتين", "منذ %d دقائق", "منذ %d دقيقة")
}

func arabicFutureMinutes(n int) string {
	return arabicPlural(n, "بعد دقيقة واحدة", "بعد دقيقتين", "بعد %d دقائق", "بعد %d دقيقة")
}

func arabicPastHours(n int) string {
	return arabicPlural(n, "منذ ساعة واحدة", "منذ ساعتين", "منذ %d ساعات", "منذ %d ساعة")
}

func arabicFutureHours(n int) string {
	return arabicPlural(n, "بعد ساعة واحدة", "بعد ساعتين", "بعد %d ساعات", "بعد %d ساعة")
}

func arabicPastDays(n int) string {
	return arabicPlural(n, "منذ يوم واحد", "منذ يومين", "منذ %d أيام", "منذ %d يومًا")
}

func arabicFutureDays(n int) string {
	return arabicPlural(n, "بعد يوم واحد", "بعد يومين", "بعد %d أيام", "بعد %d يومًا")
}

func arabicPastMonths(n int) string {
	return arabicPlural(n, "منذ شهر واحد", "منذ شهرين", "منذ %d أشهر", "منذ %d شهرًا")
}

func arabicFutureMonths(n int) string {
	return arabicPlural(n, "بعد شهر واحد", "بعد شهرين", "بعد %d أشهر", "بعد %d شهرًا")
}

func arabicPastYears(n int) string {
	return arabicPlural(n, "منذ سنة واحدة", "منذ سنتين", "منذ %d سنوات", "منذ %d سنة")
}

func arabicFutureYears(n int) string {
	return arabicPlural(n, "بعد سنة واحدة", "بعد سنتين", "بعد %d سنوات", "بعد %d سنة")
}

var arabicSaudiArabia = Formatters{
	-1 << 63: func(d time.Duration) string {
		years := int(math.Ceil(-d.Hours() / (24.0 * 30 * 12)))
		return arabicFutureYears(years)
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "بعد سنة واحدة" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		months := int(math.Ceil(-d.Hours() / (24.0 * 30)))
		return arabicFutureMonths(months)
	},
	-45 * (24 * time.Hour): func(_ time.Duration) string { return "بعد شهر واحد" },
	-25 * (24 * time.Hour): func(d time.Duration) string {
		days := int(math.Ceil(-d.Hours() / 24.0))
		return arabicFutureDays(days)
	},
	-35 * time.Hour: func(_ time.Duration) string { return "بعد يوم واحد" },
	-21 * time.Hour: func(d time.Duration) string {
		hours := int(math.Ceil(-d.Hours()))
		return arabicFutureHours(hours)
	},
	-89 * time.Minute: func(_ time.Duration) string { return "بعد ساعة واحدة" },
	-44 * time.Minute: func(d time.Duration) string {
		minutes := int(math.Ceil(-d.Minutes()))
		return arabicFutureMinutes(minutes)
	},
	-89 * time.Second: func(_ time.Duration) string { return "بعد دقيقة واحدة" },
	-44 * time.Second: func(_ time.Duration) string { return "بعد ثوانٍ قليلة" },

	44 * time.Second: func(_ time.Duration) string { return "منذ ثوانٍ قليلة" },
	89 * time.Second: func(_ time.Duration) string { return "منذ دقيقة واحدة" },
	44 * time.Minute: func(d time.Duration) string {
		minutes := int(math.Ceil(d.Minutes()))
		return arabicPastMinutes(minutes)
	},
	89 * time.Minute: func(_ time.Duration) string { return "منذ ساعة واحدة" },
	21 * time.Hour: func(d time.Duration) string {
		hours := int(math.Ceil(d.Hours()))
		return arabicPastHours(hours)
	},
	35 * time.Hour: func(_ time.Duration) string { return "منذ يوم واحد" },
	25 * (24 * time.Hour): func(d time.Duration) string {
		days := int(math.Ceil(d.Hours() / 24.0))
		return arabicPastDays(days)
	},
	45 * (24 * time.Hour): func(_ time.Duration) string { return "منذ شهر واحد" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		months := int(math.Ceil(d.Hours() / (24.0 * 30)))
		return arabicPastMonths(months)
	},
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "منذ سنة واحدة" },
	1<<63 - 1: func(d time.Duration) string {
		years := int(math.Ceil(d.Hours() / (24.0 * 30 * 12)))
		return arabicPastYears(years)
	},
}

func init() {
	Register("ar", arabicSaudiArabia)
	Register("ar-SA", arabicSaudiArabia)
}
