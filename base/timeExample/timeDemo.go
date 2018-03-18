package timeExample

import (
	"time"
)

type TimeDemo struct{}

func NewTimeDemo() *TimeDemo {
	return &TimeDemo{}
}

func (t *TimeDemo) AddDateTime(year, month, day int, format string) string {
	tempTime := time.Now()
	addTime := tempTime.AddDate(year, month, day)
	return addTime.Format(format)
}

func (t *TimeDemo) CompareTime(a string, b string) int {
	a1, _ := time.Parse("2006-01-02 15:04:05", a)
	b1, _ := time.Parse("2006-01-02 15:04:05", b)
	if a1.After(b1) {
		return 1
	} else if a1.Before(b1) {
		return -1
	} else {
		return 0
	}
}

func (t *TimeDemo) StrToUnix(format, str string) int64 {
	tempTime, _ := time.Parse(format, str)
	return tempTime.Unix()
}

func (t *TimeDemo) UnixToStr(format string, unix int64) string {
	strTime := time.Unix(unix, 0).Format(format)
	return strTime
}

func (t *TimeDemo) CurrentStrTime(format string) string {
	return time.Now().Format(format)
}

func (t *TimeDemo) CurrentUnixTime() int64 {
	return time.Now().Unix()
}
