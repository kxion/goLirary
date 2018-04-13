package timeExample

import (
	"fmt"
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

func (t *TimeDemo) GetMonth(tt time.Time) int64 {
	var month int64
	monthString := tt.Month().String()
	switch monthString {
	case "January":
		month = 1
	case "February":
		month = 2
	case "March":
		month = 3
	case "April":
		month = 4
	case "May":
		month = 5
	case "June":
		month = 6
	case "July":
		month = 7
	case "August":
		month = 8
	case "September":
		month = 9
	case "October":
		month = 10
	case "November":
		month = 11
	case "December":
		month = 12
	}
	return month
}

func (t *TimeDemo) GetLastDayOfMonth(year int64, month int64) (string, string) {
	formatMonthStr := "%d-%d"
	if month < 10 {
		formatMonthStr = "%d-0%d"
	}
	monthTimeStr := fmt.Sprintf(formatMonthStr, year, month)
	monthTime, _ := time.Parse("2006-01", monthTimeStr)
	firstDay := monthTime.AddDate(0, 0, 0).Format("2006-01-02")
	endDay := monthTime.AddDate(0, 1, -1).Format("2006-01-02")
	return firstDay, endDay
}
