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

func (t *TimeDemo) FormatHour(tt time.Time) string {
	return tt.Format("15:04")
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

func (t *TimeDemo) GetWeek(tt time.Time) int64 {
	var weekDay int64
	weekString := tt.Weekday().String()
	switch weekString {
	case "Sunday":
		weekDay = 0
	case "Monday":
		weekDay = 1
	case "Tuesday":
		weekDay = 2
	case "Wednesday":
		weekDay = 3
	case "Thursday":
		weekDay = 4
	case "Friday":
		weekDay = 5
	case "Saturday":
		weekDay = 6
	}
	return weekDay
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

func (t *TimeDemo) GetPeriodOfMonthDay(startDate, endDate string) []string {
	var periodMonth []string
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)
	startMonthInt := t.GetMonth(startTime)
	endMonthInt := t.GetMonth(endTime)
	startYearInt := startTime.Year()
	endYearInt := endTime.Year()
	if startYearInt == endYearInt {
		if startMonthInt == endMonthInt {
			startDay := startDate
			endDay := endDate
			periodMonth = append(periodMonth, startDay, endDay)
		}
		for i := startMonthInt; i <= endMonthInt; i++ {
			startDay, endDay := t.GetLastDayOfMonth(int64(startYearInt), int64(i))
			if i == startMonthInt {
				startDay = startDate
			}
			if i == endMonthInt {
				endDay = endDate
			}
			periodMonth = append(periodMonth, startDay, endDay)
		}
	} else {
		for i := startYearInt; i <= endYearInt; i++ {
			if i == startYearInt {
				for j := startMonthInt; j <= 12; j++ {
					startDay, endDay := t.GetLastDayOfMonth(int64(startYearInt), int64(j))
					if j == startMonthInt {
						startDay = startDate
					}
					periodMonth = append(periodMonth, startDay, endDay)
				}
			} else if i == endYearInt {
				for j := 1; j <= int(endMonthInt); j++ {
					startDay, endDay := t.GetLastDayOfMonth(int64(endYearInt), int64(j))
					if j == int(endMonthInt) {
						endDay = endDate
					}
					periodMonth = append(periodMonth, startDay, endDay)
				}
			} else {
				for j := 1; j <= 12; j++ {
					startDay, endDay := t.GetLastDayOfMonth(int64(i), int64(j))
					periodMonth = append(periodMonth, startDay, endDay)
				}
			}
		}
	}
	return periodMonth
}
