package locale

import (
	"fmt"
	"math"
	"time"
)

var chineseUnitedStates = Formatters{
	-1 << 63:                    func(d time.Duration) string { return fmt.Sprintf("%.0f 年后", math.Ceil(-d.Hours()/(24.0*30*12))) },
	-17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "1年后" },
	-10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f 个月后", math.Ceil(-d.Hours()/(24.0*30))) },
	-45 * (24 * time.Hour):      func(d time.Duration) string { return "1个月后" },
	-25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f 天后", math.Ceil(-d.Hours()/24.0)) },
	-35 * time.Hour:             func(_ time.Duration) string { return "1天后" },
	-21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f 小时后", math.Ceil(-d.Hours())) },
	-89 * time.Minute:           func(_ time.Duration) string { return "1小时后" },
	-44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f 分钟后", math.Ceil(-d.Minutes())) },
	-89 * time.Second:           func(_ time.Duration) string { return "1分钟后" },
	-44 * time.Second:           func(_ time.Duration) string { return "1秒后" },

	44 * time.Second:           func(_ time.Duration) string { return "1秒前" },
	89 * time.Second:           func(_ time.Duration) string { return "1分钟前" },
	44 * time.Minute:           func(d time.Duration) string { return fmt.Sprintf("%.0f 分钟前", math.Ceil(d.Minutes())) },
	89 * time.Minute:           func(_ time.Duration) string { return "1小时前" },
	21 * time.Hour:             func(d time.Duration) string { return fmt.Sprintf("%.0f 小时前", math.Ceil(d.Hours())) },
	35 * time.Hour:             func(_ time.Duration) string { return "1天前" },
	25 * (24 * time.Hour):      func(d time.Duration) string { return fmt.Sprintf("%.0f 天前", math.Ceil(d.Hours()/24.0)) },
	45 * (24 * time.Hour):      func(d time.Duration) string { return "1个月前" },
	10 * (24 * time.Hour) * 30: func(d time.Duration) string { return fmt.Sprintf("%.0f 个月前", math.Ceil(d.Hours()/(24.0*30))) },
	17 * (24 * time.Hour) * 30: func(_ time.Duration) string { return "1年前" },
	1<<63 - 1:                  func(d time.Duration) string { return fmt.Sprintf("%.0f 年前", math.Ceil(d.Hours()/(24.0*30*12))) },
}

func init() {
	Register("zh-CN", chineseUnitedStates)
}
