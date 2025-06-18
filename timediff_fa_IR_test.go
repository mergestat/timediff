package timediff_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/mergestat/timediff"
	"github.com/mergestat/timediff/locale"
)

var fixtures_fa_IR = map[string]string{
	"-10s":                            "چند ثانیه پیش",
	"-44s":                            "چند ثانیه پیش",
	"-45s":                            "یک دقیقه پیش",
	"-89s":                            "یک دقیقه پیش",
	"-90s":                            "2 دقیقه پیش",
	"-91s":                            "2 دقیقه پیش",
	"-2m":                             "2 دقیقه پیش",
	"-10m":                            "10 دقیقه پیش",
	"-44m":                            "44 دقیقه پیش",
	"-45m":                            "یک ساعت پیش",
	"-60m":                            "یک ساعت پیش",
	"-1h":                             "یک ساعت پیش",
	"-80m":                            "یک ساعت پیش",
	"-89m":                            "یک ساعت پیش",
	"-90m":                            "2 ساعت پیش",
	"-2h":                             "2 ساعت پیش",
	"-20h":                            "20 ساعت پیش",
	"-21h":                            "21 ساعت پیش",
	"-21h30m":                         "یک روز پیش",
	"-22h":                            "یک روز پیش",
	"-24h":                            "یک روز پیش",
	"-24h30m":                         "یک روز پیش",
	"-34h59m":                         "یک روز پیش",
	"-36h":                            "2 روز پیش",
	fmt.Sprintf("-%dh", 10*24):        "10 روز پیش",
	fmt.Sprintf("-%dh", 25*24):        "25 روز پیش",
	fmt.Sprintf("-%dh", 26*24):        "یک ماه پیش",
	fmt.Sprintf("-%dh", 45*24):        "یک ماه پیش",
	fmt.Sprintf("-%dh2m", 45*24):      "2 ماه پیش",
	fmt.Sprintf("-%dh", 46*24+1):      "2 ماه پیش",
	fmt.Sprintf("-%dh", 80*24):        "3 ماه پیش",
	fmt.Sprintf("-%dh", 9*24*30):      "9 ماه پیش",
	fmt.Sprintf("-%dh", 10*24*30):     "10 ماه پیش",
	fmt.Sprintf("-%dh1m", 10*24*30):   "یک سال پیش",
	fmt.Sprintf("-%dh", 12*24*30):     "یک سال پیش",
	fmt.Sprintf("-%dh", 17*24*30+1):   "2 سال پیش",
	fmt.Sprintf("-%dh", 24*24*30):     "2 سال پیش",
	fmt.Sprintf("-%dh", 20*24*30*12):  "20 سال پیش",
	fmt.Sprintf("-%dh", 100*24*30*12): "100 سال پیش",

	"10s":                            "در چند ثانیه",
	"44s":                            "در چند ثانیه",
	"45s":                            "در یک دقیقه",
	"89s":                            "در یک دقیقه",
	"90s":                            "در 2 دقیقه",
	"2m":                             "در 2 دقیقه",
	"10m":                            "در 10 دقیقه",
	"44m":                            "در 44 دقیقه",
	"45m":                            "در یک ساعت",
	"60m":                            "در یک ساعت",
	"1h":                             "در یک ساعت",
	"80m":                            "در یک ساعت",
	"89m":                            "در یک ساعت",
	"89m10s":                         "در 2 ساعت",
	"90m":                            "در 2 ساعت",
	"2h":                             "در 2 ساعت",
	"20h":                            "در 20 ساعت",
	"21h":                            "در 21 ساعت",
	"21h30m":                         "در یک روز",
	"22h":                            "در یک روز",
	"24h":                            "در یک روز",
	"36h":                            "در 2 روز",
	"35h10m":                         "در 2 روز",
	fmt.Sprintf("%dh", 10*24):        "در 10 روز",
	fmt.Sprintf("%dh", 25*24):        "در 25 روز",
	fmt.Sprintf("%dh", 26*24):        "در یک ماه",
	fmt.Sprintf("%dh", 45*24):        "در یک ماه",
	fmt.Sprintf("%dh1m", 45*24):      "در 2 ماه",
	fmt.Sprintf("%dh", 46*24):        "در 2 ماه",
	fmt.Sprintf("%dh", 80*24):        "در 3 ماه",
	fmt.Sprintf("%dh", 9*24*30):      "در 9 ماه",
	fmt.Sprintf("%dh", 10*24*30):     "در 10 ماه",
	fmt.Sprintf("%dh1m", 10*24*30):   "در یک سال",
	fmt.Sprintf("%dh", 12*24*30):     "در یک سال",
	fmt.Sprintf("%dh", 24*24*30):     "در 2 سال",
	fmt.Sprintf("%dh", 20*24*30*12):  "در 20 سال",
	fmt.Sprintf("%dh", 100*24*30*12): "در 100 سال",
}

func TestTimeDiffFaIR(t *testing.T) {
	execFixtures(t, fixtures_fa_IR, timediff.WithLocale("fa-IR"))
}

func TestWithStartTimeFaIR(t *testing.T) {
	start := time.Date(2022, time.January, 22, 12, 0, 0, 0, time.Now().Local().Location())
	timeToDiff := time.Date(2022, time.January, 22, 10, 0, 0, 0, time.Now().Local().Location())

	want := "2 ساعت پیش"
	got := timediff.TimeDiff(timeToDiff, timediff.WithLocale("fa-IR"), timediff.WithStartTime(start))

	if got != want {
		t.Fatalf("expected: %q, got %q", want, got)
	}
}

func TestWithCustomFormattersFaIR(t *testing.T) {
	var custom = locale.Formatters{
		21 * time.Hour: func(d time.Duration) string { return fmt.Sprintf("پیام: %.0f ساعت", math.Ceil(d.Hours())) },
	}

	{
		want := "پیام: 2 ساعت"
		got := timediff.TimeDiff(time.Now().Add(-2*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("fa-IR"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}

	{
		want := "پیام: 100 ساعت"
		got := timediff.TimeDiff(time.Now().Add(-100*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("fa-IR"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}
}
