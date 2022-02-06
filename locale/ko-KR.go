package locale

import (
	"fmt"
	"math"
	"time"
)

var koreanKorea = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("%.0f년 뒤", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "내년" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f개월 뒤", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "다음 달" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f일 뒤", math.Ceil(-d.Hours()/24.0)) },
	-59 * time.Hour:             func(_ time.Duration) string { return "모레" },
	-35 * time.Hour:             func(_ time.Duration) string { return "내일" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f시간 뒤", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "약 한 시간 뒤" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f분 뒤", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "약 1분 뒤" },
	-44 * time.Second:           func(_ time.Duration) string { return "곧" },

	44 * time.Second:           func(_ time.Duration) string { return "방금" },
	89 * time.Second:           func(_ time.Duration) string { return "약 1분 전" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f분 전", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "약 한 시간 전" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f시간 전", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "어제" },
	59 * time.Hour:             func(_ time.Duration) string { return "그제" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f일 전", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "저번 달" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f개월 전", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "작년" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f년 전", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("ko", koreanKorea)
	Register("ko-KR", koreanKorea)
}
