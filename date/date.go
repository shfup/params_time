package date

import "time"

const (
	DayFormat = "2006-01-02"
)

// 开始日期和结束日期转换为时间
// style:0-start,1-end
func DayToStartEndTime(t string, style int) string {
	_, err := time.Parse(DayFormat, t)
	if err != nil {
		return ""
	}
	if style == 0 {
		t += " 00:00:00"
	} else {
		t += " 23:59:59"
	}
	return t
}
