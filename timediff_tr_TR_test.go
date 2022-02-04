package timediff_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/mergestat/timediff"
	"github.com/mergestat/timediff/locale"
)

var fixtures_tr_TR = map[string]string{
	"-10s":                            "birkaç saniye önce",
	"-44s":                            "birkaç saniye önce",
	"-45s":                            "bir dakika önce",
	"-89s":                            "bir dakika önce",
	"-90s":                            "2 dakika önce",
	"-91s":                            "2 dakika önce",
	"-2m":                             "2 dakika önce",
	"-10m":                            "10 dakika önce",
	"-44m":                            "44 dakika önce",
	"-45m":                            "bir saat önce",
	"-60m":                            "bir saat önce",
	"-1h":                             "bir saat önce",
	"-80m":                            "bir saat önce",
	"-89m":                            "bir saat önce",
	"-90m":                            "2 saat önce",
	"-2h":                             "2 saat önce",
	"-20h":                            "20 saat önce",
	"-21h":                            "21 saat önce",
	"-21h30m":                         "bir gün önce",
	"-22h":                            "bir gün önce",
	"-24h":                            "bir gün önce",
	"-24h30m":                         "bir gün önce",
	"-34h59m":                         "bir gün önce",
	"-36h":                            "2 gün önce",
	fmt.Sprintf("-%dh", 10*24):        "10 gün önce",
	fmt.Sprintf("-%dh", 25*24):        "25 gün önce",
	fmt.Sprintf("-%dh", 26*24):        "bir ay önce",
	fmt.Sprintf("-%dh", 45*24):        "bir ay önce",
	fmt.Sprintf("-%dh2m", 45*24):      "2 ay önce",
	fmt.Sprintf("-%dh", 46*24+1):      "2 ay önce",
	fmt.Sprintf("-%dh", 80*24):        "3 ay önce",
	fmt.Sprintf("-%dh", 9*24*30):      "9 ay önce",
	fmt.Sprintf("-%dh", 10*24*30):     "10 ay önce",
	fmt.Sprintf("-%dh1m", 10*24*30):   "bir yıl önce",
	fmt.Sprintf("-%dh", 12*24*30):     "bir yıl önce",
	fmt.Sprintf("-%dh", 17*24*30+1):   "2 yıl önce",
	fmt.Sprintf("-%dh", 24*24*30):     "2 yıl önce",
	fmt.Sprintf("-%dh", 20*24*30*12):  "20 yıl önce",
	fmt.Sprintf("-%dh", 100*24*30*12): "100 yıl önce",

	"10s":                            "birkaç saniye içinde",
	"44s":                            "birkaç saniye içinde",
	"45s":                            "bir dakika içinde",
	"89s":                            "bir dakika içinde",
	"90s":                            "2 dakika içinde",
	"2m":                             "2 dakika içinde",
	"10m":                            "10 dakika içinde",
	"44m":                            "44 dakika içinde",
	"45m":                            "bir saat içinde",
	"60m":                            "bir saat içinde",
	"1h":                             "bir saat içinde",
	"80m":                            "bir saat içinde",
	"89m":                            "bir saat içinde",
	"89m10s":                         "2 saat içinde",
	"90m":                            "2 saat içinde",
	"2h":                             "2 saat içinde",
	"20h":                            "20 saat içinde",
	"21h":                            "21 saat içinde",
	"21h30m":                         "bir gün içinde",
	"22h":                            "bir gün içinde",
	"24h":                            "bir gün içinde",
	"35h10m":                         "2 gün içinde",
	"36h":                            "2 gün içinde",
	fmt.Sprintf("%dh", 10*24):        "10 gün içinde",
	fmt.Sprintf("%dh", 25*24):        "25 gün içinde",
	fmt.Sprintf("%dh", 26*24):        "bir ay içinde",
	fmt.Sprintf("%dh", 45*24):        "bir ay içinde",
	fmt.Sprintf("%dh1m", 45*24):      "2 ay içinde",
	fmt.Sprintf("%dh", 46*24):        "2 ay içinde",
	fmt.Sprintf("%dh", 80*24):        "3 ay içinde",
	fmt.Sprintf("%dh", 9*24*30):      "9 ay içinde",
	fmt.Sprintf("%dh", 10*24*30):     "10 ay içinde",
	fmt.Sprintf("%dh1m", 10*24*30):   "bir yıl içinde",
	fmt.Sprintf("%dh", 12*24*30):     "bir yıl içinde",
	fmt.Sprintf("%dh", 24*24*30):     "2 yıl içinde",
	fmt.Sprintf("%dh", 20*24*30*12):  "20 yıl içinde",
	fmt.Sprintf("%dh", 100*24*30*12): "100 yıl içinde",
}

func TestTimeDiffTrTR(t *testing.T) {
	execFixtures(t, fixtures_tr_TR, timediff.WithLocale("tr-TR"))
}

func TestWithStartTimeTrTR(t *testing.T) {
	start := time.Date(2022, time.January, 22, 12, 0, 0, 0, time.Now().Local().Location())
	timeToDiff := time.Date(2022, time.January, 22, 10, 0, 0, 0, time.Now().Local().Location())

	want := "2 saat önce"
	got := timediff.TimeDiff(timeToDiff, timediff.WithLocale("tr-TR"), timediff.WithStartTime(start))

	if got != want {
		t.Fatalf("expected: %q, got %q", want, got)
	}
}

func TestWithCustomFormattersTrTR(t *testing.T) {
	var custom = locale.Formatters{
		21 * time.Hour: func(d time.Duration) string { return fmt.Sprintf("Mesaj: %.0f saat", math.Ceil(d.Hours())) },
	}

	{
		want := "Mesaj: 2 saat"
		got := timediff.TimeDiff(time.Now().Add(-2*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("tr-TR"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}

	{
		want := "Mesaj: 100 saat"
		got := timediff.TimeDiff(time.Now().Add(-100*time.Hour), timediff.WithCustomFormatters(custom), timediff.WithLocale("tr-TR"))

		if got != want {
			t.Fatalf("expected: %q, got %q", want, got)
		}
	}
}
