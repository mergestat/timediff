package locale

import (
	"fmt"
	"math"
	"time"
)

var portugueseBrazil = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("em %.0f anos", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "em um ano" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("em %.0f meses", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "em um mês" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("em %.0f dias", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "em um dia" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("em %.0f horas", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "em uma hora" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("em %.0f minutos", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "em um minuto" },
	-44 * time.Second:           func(_ time.Duration) string { return "em poucos segundos" },

	44 * time.Second:           func(_ time.Duration) string { return "há poucos segundos" },
	89 * time.Second:           func(_ time.Duration) string { return "há um minuto" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f minutos atrás", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "há uma hora" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f horas atrás", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "há um dia" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f dias atrás", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "há um mês" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f meses atrás", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "há um ano" },
	1<<63 - 1: func(d time.Duration) string {
		return fmt.Sprintf("%.0f anos atrás", math.Ceil(d.Hours()/(24.0*30*12)))
	},
}

func init() {
	Register("pt", portugueseBrazil) // also register it as the default portuguese locale
	Register("pt-BR", portugueseBrazil)
}
