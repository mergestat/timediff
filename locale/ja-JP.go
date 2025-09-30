package locale

import (
	"fmt"
	"math"
	"time"
)

var japaneseJapan = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("%.0f年後", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "1年後" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0fヶ月後", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "1ヶ月後" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f日後", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "1日後" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f時間後", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "1時間後" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f分後", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "1分後" },
	-44 * time.Second:           func(_ time.Duration) string { return "数秒後" },

	44 * time.Second:           func(_ time.Duration) string { return "数秒前" },
	89 * time.Second:           func(_ time.Duration) string { return "1分前" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f分前", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "1時間前" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f時間前", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "1日前" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f日前", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "1ヶ月前" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0fヶ月前", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "1年前" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f年前", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("ja", japaneseJapan)
	Register("ja-JP", japaneseJapan)
}