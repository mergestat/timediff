package timediff_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/mergestat/timediff"
	"github.com/mergestat/timediff/locale"
)

var fixtures_ar_SA = map[string]string{
	"-10s":                            "منذ ثوانٍ قليلة",
	"-44s":                            "منذ ثوانٍ قليلة",
	"-45s":                            "منذ دقيقة واحدة",
	"-89s":                            "منذ دقيقة واحدة",
	"-90s":                            "منذ دقيقتين",
	"-91s":                            "منذ دقيقتين",
	"-2m":                             "منذ دقيقتين",
	"-10m":                            "منذ 10 دقائق",
	"-44m":                            "منذ 44 دقيقة",
	"-45m":                            "منذ ساعة واحدة",
	"-60m":                            "منذ ساعة واحدة",
	"-1h":                             "منذ ساعة واحدة",
	"-80m":                            "منذ ساعة واحدة",
	"-89m":                            "منذ ساعة واحدة",
	"-90m":                            "منذ ساعتين",
	"-2h":                             "منذ ساعتين",
	"-20h":                            "منذ 20 ساعة",
	"-21h":                            "منذ 21 ساعة",
	"-21h30m":                         "منذ يوم واحد",
	"-22h":                            "منذ يوم واحد",
	"-24h":                            "منذ يوم واحد",
	"-24h30m":                         "منذ يوم واحد",
	"-34h59m":                         "منذ يوم واحد",
	"-36h":                            "منذ يومين",
	fmt.Sprintf("-%dh", 10*24):        "منذ 10 أيام",
	fmt.Sprintf("-%dh", 25*24):        "منذ 25 يومًا",
	fmt.Sprintf("-%dh", 26*24):        "منذ شهر واحد",
	fmt.Sprintf("-%dh", 45*24):        "منذ شهر واحد",
	fmt.Sprintf("-%dh2m", 45*24):      "منذ شهرين",
	fmt.Sprintf("-%dh", 46*24+1):      "منذ شهرين",
	fmt.Sprintf("-%dh", 80*24):        "منذ 3 أشهر",
	fmt.Sprintf("-%dh", 9*24*30):      "منذ 9 أشهر",
	fmt.Sprintf("-%dh", 10*24*30):     "منذ 10 أشهر",
	fmt.Sprintf("-%dh1m", 10*24*30):   "منذ سنة واحدة",
	fmt.Sprintf("-%dh", 12*24*30):     "منذ سنة واحدة",
	fmt.Sprintf("-%dh", 17*24*30+1):   "منذ سنتين",
	fmt.Sprintf("-%dh", 24*24*30):     "منذ سنتين",
	fmt.Sprintf("-%dh", 20*24*30*12):  "منذ 20 سنة",
	fmt.Sprintf("-%dh", 100*24*30*12): "منذ 100 سنة",

	"10s":                            "بعد ثوانٍ قليلة",
	"44s":                            "بعد ثوانٍ قليلة",
	"45s":                            "بعد دقيقة واحدة",
	"89s":                            "بعد دقيقة واحدة",
	"90s":                            "بعد دقيقتين",
	"2m":                             "بعد دقيقتين",
	"10m":                            "بعد 10 دقائق",
	"44m":                            "بعد 44 دقيقة",
	"45m":                            "بعد ساعة واحدة",
	"60m":                            "بعد ساعة واحدة",
	"1h":                             "بعد ساعة واحدة",
	"80m":                            "بعد ساعة واحدة",
	"89m":                            "بعد ساعة واحدة",
	"89m10s":                         "بعد ساعتين",
	"90m":                            "بعد ساعتين",
	"2h":                             "بعد ساعتين",
	"20h":                            "بعد 20 ساعة",
	"21h":                            "بعد 21 ساعة",
	"21h30m":                         "بعد يوم واحد",
	"22h":                            "بعد يوم واحد",
	"24h":                            "بعد يوم واحد",
	"35h10m":                         "بعد يومين",
	"36h":                            "بعد يومين",
	fmt.Sprintf("%dh", 10*24):        "بعد 10 أيام",
	fmt.Sprintf("%dh", 25*24):        "بعد 25 يومًا",
	fmt.Sprintf("%dh", 26*24):        "بعد شهر واحد",
	fmt.Sprintf("%dh", 45*24):        "بعد شهر واحد",
	fmt.Sprintf("%dh1m", 45*24):      "بعد شهرين",
	fmt.Sprintf("%dh", 46*24):        "بعد شهرين",
	fmt.Sprintf("%dh", 80*24):        "بعد 3 أشهر",
	fmt.Sprintf("%dh", 9*24*30):      "بعد 9 أشهر",
	fmt.Sprintf("%dh", 10*24*30):     "بعد 10 أشهر",
	fmt.Sprintf("%dh1m", 10*24*30):   "بعد سنة واحدة",
	fmt.Sprintf("%dh", 12*24*30):     "بعد سنة واحدة",
	fmt.Sprintf("%dh", 24*24*30):     "بعد سنتين",
	fmt.Sprintf("%dh", 20*24*30*12):  "بعد 20 سنة",
	fmt.Sprintf("%dh", 100*24*30*12): "بعد 100 سنة",
}

func TestTimeDiffArSA(t *testing.T) {
	execFixtures(t, fixtures_ar_SA, timediff.WithLocale("ar-SA"))
}

func TestWithStartTimeArSA(t *testing.T) {
	start := time.Date(2022, time.January, 22, 12, 0, 0, 0, time.Now().Local().Location())
	timeToDiff := time.Date(2022, time.January, 22, 10, 0, 0, 0, time.Now().Local().Location())

	want := "منذ ساعتين"
	got := timediff.TimeDiff(timeToDiff, timediff.WithLocale("ar-SA"), timediff.WithStartTime(start))

	if got != want {
		t.Fatalf("expected: %q, got %q", want, got)
	}
}

func TestWithCustomFormattersArSA(t *testing.T) {
	var custom = locale.Formatters{
		21 * time.Hour: func(d time.Duration) string {
			return fmt.Sprintf("رسالة مخصصة: %.0f ساعة", math.Ceil(d.Hours()))
		},
	}

	{
		want := "رسالة مخصصة: 2 ساعة"
		got := timediff.TimeDiff(time.Now().Add(-2*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("ar-SA"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}

	{
		want := "رسالة مخصصة: 100 ساعة"
		got := timediff.TimeDiff(time.Now().Add(-100*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("ar-SA"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}
}
