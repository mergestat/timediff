package timediff_test

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

var fixtures_ptBR = map[string]string{
	"-10s":                            "há poucos segundos",
	"-44s":                            "há poucos segundos",
	"-45s":                            "há um minuto",
	"-89s":                            "há um minuto",
	"-90s":                            "2 minutos atrás",
	"-91s":                            "2 minutos atrás",
	"-2m":                             "2 minutos atrás",
	"-10m":                            "10 minutos atrás",
	"-44m":                            "44 minutos atrás",
	"-45m":                            "há uma hora",
	"-60m":                            "há uma hora",
	"-1h":                             "há uma hora",
	"-80m":                            "há uma hora",
	"-89m":                            "há uma hora",
	"-90m":                            "2 horas atrás",
	"-2h":                             "2 horas atrás",
	"-20h":                            "20 horas atrás",
	"-21h":                            "21 horas atrás",
	"-21h30m":                         "há um dia",
	"-22h":                            "há um dia",
	"-24h":                            "há um dia",
	"-24h30m":                         "há um dia",
	"-34h59m":                         "há um dia",
	"-36h":                            "2 dias atrás",
	fmt.Sprintf("-%dh", 10*24):        "10 dias atrás",
	fmt.Sprintf("-%dh", 25*24):        "25 dias atrás",
	fmt.Sprintf("-%dh", 26*24):        "há um mês",
	fmt.Sprintf("-%dh", 45*24):        "há um mês",
	fmt.Sprintf("-%dh2m", 45*24):      "2 meses atrás",
	fmt.Sprintf("-%dh", 46*24+1):      "2 meses atrás",
	fmt.Sprintf("-%dh", 80*24):        "3 meses atrás",
	fmt.Sprintf("-%dh", 9*24*30):      "9 meses atrás",
	fmt.Sprintf("-%dh", 10*24*30):     "10 meses atrás",
	fmt.Sprintf("-%dh1m", 10*24*30):   "há um ano",
	fmt.Sprintf("-%dh", 12*24*30):     "há um ano",
	fmt.Sprintf("-%dh", 17*24*30+1):   "2 anos atrás",
	fmt.Sprintf("-%dh", 24*24*30):     "2 anos atrás",
	fmt.Sprintf("-%dh", 20*24*30*12):  "20 anos atrás",
	fmt.Sprintf("-%dh", 100*24*30*12): "100 anos atrás",

	"10s":                            "em poucos segundos",
	"44s":                            "em poucos segundos",
	"45s":                            "em um minuto",
	"89s":                            "em um minuto",
	"90s":                            "em 2 minutos",
	"2m":                             "em 2 minutos",
	"10m":                            "em 10 minutos",
	"44m":                            "em 44 minutos",
	"45m":                            "em uma hora",
	"60m":                            "em uma hora",
	"1h":                             "em uma hora",
	"80m":                            "em uma hora",
	"89m":                            "em uma hora",
	"89m10s":                         "em 2 horas",
	"90m":                            "em 2 horas",
	"2h":                             "em 2 horas",
	"20h":                            "em 20 horas",
	"21h":                            "em 21 horas",
	"21h30m":                         "em um dia",
	"22h":                            "em um dia",
	"24h":                            "em um dia",
	"35h10m":                         "em 2 dias",
	"36h":                            "em 2 dias",
	fmt.Sprintf("%dh", 10*24):        "em 10 dias",
	fmt.Sprintf("%dh", 25*24):        "em 25 dias",
	fmt.Sprintf("%dh", 26*24):        "em um mês",
	fmt.Sprintf("%dh", 45*24):        "em um mês",
	fmt.Sprintf("%dh1m", 45*24):      "em 2 meses",
	fmt.Sprintf("%dh", 46*24):        "em 2 meses",
	fmt.Sprintf("%dh", 80*24):        "em 3 meses",
	fmt.Sprintf("%dh", 9*24*30):      "em 9 meses",
	fmt.Sprintf("%dh", 10*24*30):     "em 10 meses",
	fmt.Sprintf("%dh1m", 10*24*30):   "em um ano",
	fmt.Sprintf("%dh", 12*24*30):     "em um ano",
	fmt.Sprintf("%dh", 24*24*30):     "em 2 anos",
	fmt.Sprintf("%dh", 20*24*30*12):  "em 20 anos",
	fmt.Sprintf("%dh", 100*24*30*12): "em 100 anos",
}

func TestTimeDiffPtBR(t *testing.T) {
	now := time.Now()

	durations := make([]string, 0, len(fixtures_ptBR))
	wants := make([]string, len(fixtures_ptBR))

	for d := range fixtures_ptBR {
		durations = append(durations, d)
	}

	sort.SliceStable(durations, func(i, j int) bool {
		pi, err := time.ParseDuration(durations[i])
		if err != nil {
			t.Fatal(err)
		}
		pj, err := time.ParseDuration(durations[j])
		if err != nil {
			t.Fatal(err)
		}
		return pi < pj
	})

	// populate the slice of formatters, corresponding to their sorted durations
	for i, d := range durations {
		wants[i] = fixtures_ptBR[d]
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)

	for d, durStr := range durations {
		want := wants[d]
		dur, err := time.ParseDuration(durStr)
		if err != nil {
			t.Fatal(err)
		}

		timeToDiff := now.Add(dur)
		got := timediff.TimeDiff(timeToDiff, timediff.WithLocale("pt-BR"))

		if got != want {
			t.Fatalf("expected: %q, got: %q for duration: %q (%q)", want, got, durStr, dur)
		}

		fmt.Fprintln(w, strings.Join([]string{durStr, got}, "\t"))
	}

	if err := w.Flush(); err != nil {
		t.Fatal(err)
	}
}

func TestWithStartTimePtBR(t *testing.T) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.TabIndent)

	start := time.Date(2022, time.January, 22, 12, 0, 0, 0, time.Now().Local().Location())

	// past time
	timeToDiff := time.Date(2022, time.January, 22, 10, 0, 0, 0, time.Now().Local().Location())
	want := "2 horas atrás"
	got := timediff.TimeDiff(timeToDiff, timediff.WithLocale("pt-BR"), timediff.WithStartTime(start))

	if got != want {
		t.Fatalf("expected: %q, got %q", want, got)
	}
	fmt.Fprintln(w, strings.Join([]string{"-2h", got}, "\t"))

	now := time.Now()
	want = "3 minutos atrás"
	got = timediff.TimeDiff(now.Add(-3*time.Minute), timediff.WithLocale("pt-BR"), timediff.WithStartTime(now))
	if got != want {
		t.Fatalf("expected: %q, got: %q", want, got)
	}
	fmt.Fprintln(w, strings.Join([]string{"-3m", got}, "\t"))

	// future time
	timeToDiff = time.Date(2022, time.January, 22, 14, 0, 0, 0, time.Now().Local().Location())
	want = "em 2 horas"
	got = timediff.TimeDiff(timeToDiff, timediff.WithLocale("pt-BR"), timediff.WithStartTime(start))

	if got != want {
		t.Fatalf("expected: %q, got %q", want, got)
	}
	fmt.Fprintln(w, strings.Join([]string{"2h", got}, "\t"))
	if err := w.Flush(); err != nil {
		t.Fatal(err)
	}
}
