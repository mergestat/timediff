package locale

import (
	"math"
	"strings"
	"time"
)

var bengaliDigits = []rune{'০', '১', '২', '৩', '৪', '৫', '৬', '৭', '৮', '৯'}

func bengaliNumber(value float64) string {
	n := int(math.Ceil(value))
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return string(bengaliDigits[0])
	}

	var b strings.Builder

	digits := make([]rune, 0, 20)
	for n > 0 {
		digits = append(digits, bengaliDigits[n%10])
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		b.WriteRune(digits[i])
	}

	return b.String()
}

var bengaliBangladesh = Formatters{
	-1 << 63: func(d time.Duration) string {
		return bengaliNumber(-d.Hours()/(24.0*30*12)) + " বছর পরে"
	},
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "১ বছর পরে" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return bengaliNumber(-d.Hours()/(24.0*30)) + " মাস পরে"
	},
	-45 * (24 * time.Hour): func(_ time.Duration) string { return "১ মাস পরে" },
	-25 * (24 * time.Hour): func(d time.Duration) string {
		return bengaliNumber(-d.Hours()/24.0) + " দিন পরে"
	},
	-35 * time.Hour: func(_ time.Duration) string { return "১ দিন পরে" },
	-21 * time.Hour: func(d time.Duration) string {
		return bengaliNumber(-d.Hours()) + " ঘণ্টা পরে"
	},
	-89 * time.Minute: func(_ time.Duration) string { return "১ ঘণ্টা পরে" },
	-44 * time.Minute: func(d time.Duration) string {
		return bengaliNumber(-d.Minutes()) + " মিনিট পরে"
	},
	-89 * time.Second: func(_ time.Duration) string { return "১ মিনিট পরে" },
	-44 * time.Second: func(_ time.Duration) string { return "কয়েক সেকেন্ড পরে" },

	44 * time.Second: func(_ time.Duration) string { return "কয়েক সেকেন্ড আগে" },
	89 * time.Second: func(_ time.Duration) string { return "১ মিনিট আগে" },
	44 * time.Minute: func(d time.Duration) string {
		return bengaliNumber(d.Minutes()) + " মিনিট আগে"
	},
	89 * time.Minute: func(_ time.Duration) string { return "১ ঘণ্টা আগে" },
	21 * time.Hour: func(d time.Duration) string {
		return bengaliNumber(d.Hours()) + " ঘণ্টা আগে"
	},
	35 * time.Hour: func(_ time.Duration) string { return "১ দিন আগে" },
	25 * (24 * time.Hour): func(d time.Duration) string {
		return bengaliNumber(d.Hours()/24.0) + " দিন আগে"
	},
	45 * (24 * time.Hour): func(_ time.Duration) string { return "১ মাস আগে" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string {
		return bengaliNumber(d.Hours()/(24.0*30)) + " মাস আগে"
	},
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "১ বছর আগে" },
	1<<63 - 1: func(d time.Duration) string {
		return bengaliNumber(d.Hours()/(24.0*30*12)) + " বছর আগে"
	},
}

func init() {
	Register("bn", bengaliBangladesh)
	Register("bn-BD", bengaliBangladesh)
}
