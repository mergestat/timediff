package locale

import (
	"fmt"
	"math"
	"time"
)

var russianRussianFederation = Formatters{
	-1 << 63: func(d time.Duration) string {
		return fmt.Sprintf("через %.0f лет", math.Ceil(-d.Hours()/(24.0*30*12)))
	},
	-50 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("через %.0f года", math.Ceil(-d.Hours()/(24.0*30*12)))
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "через год" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("через %.0f месяцев", math.Ceil(-d.Hours()/(24.0*30)))
	},
	-4 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("через %.0f месяца", math.Ceil(-d.Hours()/(24.0*30)))
	},
	-45 * (24 * time.Hour): func(d time.Duration) string { return "через месяц" },
	-25 * (24 * time.Hour): func(d time.Duration) string {
		return fmt.Sprintf("через %.0f дней", math.Ceil(-d.Hours()/24.0))
	},
	-4 * (24 * time.Hour): func(d time.Duration) string { return fmt.Sprintf("через %.0f дня", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:       func(_ time.Duration) string { return "через день" },
	-21 * time.Hour:       func(d time.Duration) string { return fmt.Sprintf("через %.0f час", math.Ceil(-d.Hours())) },
	-20 * time.Hour:       func(d time.Duration) string { return fmt.Sprintf("через %.0f часов", math.Ceil(-d.Hours())) },
	-4 * time.Hour:        func(d time.Duration) string { return fmt.Sprintf("через %.0f часа", math.Ceil(-d.Hours())) },
	-89 * time.Minute:     func(_ time.Duration) string { return "через час" },
	-44 * time.Minute: func(d time.Duration) string {
		return fmt.Sprintf("через %.0f минуты", math.Ceil(-d.Minutes()))
	},
	-21 * time.Minute: func(d time.Duration) string {
		return fmt.Sprintf("через %.0f минут", math.Ceil(-d.Minutes()))
	},
	-4 * time.Minute: func(d time.Duration) string {
		return fmt.Sprintf("через %.0f минуты", math.Ceil(-d.Minutes()))
	},
	-89 * time.Second: func(_ time.Duration) string { return "через минуту" },
	-44 * time.Second: func(_ time.Duration) string { return "через несколько секунд" },

	44 * time.Second: func(_ time.Duration) string { return "несколько секунд назад" },
	89 * time.Second: func(_ time.Duration) string { return "минуту назад" },
	4 * time.Minute: func(d time.Duration) string {
		return fmt.Sprintf("%.0f минуты назад", math.Ceil(d.Minutes()))
	},
	44 * time.Minute:     func(d time.Duration) string { return fmt.Sprintf("%.0f минут назад", math.Ceil(d.Minutes())) },
	89 * time.Minute:     func(_ time.Duration) string { return "час назад" },
	4 * time.Hour:        func(d time.Duration) string { return fmt.Sprintf("%.0f часа назад", math.Ceil(d.Hours())) },
	20 * time.Hour:       func(d time.Duration) string { return fmt.Sprintf("%.0f часов назад", math.Ceil(d.Hours())) },
	21 * time.Hour:       func(d time.Duration) string { return fmt.Sprintf("%.0f час назад", math.Ceil(d.Hours())) },
	35 * time.Hour:       func(_ time.Duration) string { return "день назад" },
	4 * (24 * time.Hour): func(d time.Duration) string { return fmt.Sprintf("%.0f дня назад", math.Ceil(d.Hours()/24.0)) },
	25 * (24 * time.Hour): func(d time.Duration) string {
		return fmt.Sprintf("%.0f дней назад", math.Ceil(d.Hours()/24.0))
	},
	45 * (24 * time.Hour): func(d time.Duration) string { return "месяц назад" },
	4 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("%.0f месяца назад", math.Ceil(d.Hours()/(24.0*30)))
	},
	10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("%.0f месяцев назад", math.Ceil(d.Hours()/(24.0*30)))
	},
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "год назад" },
	34 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return fmt.Sprintf("%.0f года назад", math.Ceil(d.Hours()/(24.0*30*12)))
	},
	1<<63 - 1: func(d time.Duration) string {
		return fmt.Sprintf("%.0f лет назад", math.Ceil(d.Hours()/(24.0*30*12)))
	},
}

func init() {
	Register("ru", russianRussianFederation) // also register it as the default russian locale
	Register("ru-RU", russianRussianFederation)
}
