package main

import (
	"fmt"
	"goLirary/base/timeExample"
	"time"
)

func main() {
	t := timeExample.NewTimeDemo()
	startDate := "2017-05-11"
	endDate := "2019-11-15"
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
			fmt.Println(startDay, endDay)
		}
		for i := startMonthInt; i <= endMonthInt; i++ {
			startDay, endDay := t.GetLastDayOfMonth(int64(startYearInt), int64(i))
			if i == startMonthInt {
				startDay = startDate
			}
			if i == endMonthInt {
				endDay = endDate
			}
			fmt.Println(startDay, endDay)
		}
	} else {
		for i := startYearInt; i <= endYearInt; i++ {
			if i == startYearInt {
				for j := startMonthInt; j <= 12; j++ {
					startDay, endDay := t.GetLastDayOfMonth(int64(startYearInt), int64(j))
					if j == startMonthInt {
						startDay = startDate
					}
					fmt.Println(startDay, endDay)
				}
			} else if i == endYearInt {
				for j := 1; j <= int(endMonthInt); j++ {
					startDay, endDay := t.GetLastDayOfMonth(int64(endYearInt), int64(j))
					if j == int(endMonthInt) {
						endDay = endDate
					}
					fmt.Println(startDay, endDay)
				}
			} else {
				for j := 1; j <= 12; j++ {
					startDay, endDay := t.GetLastDayOfMonth(int64(i), int64(j))
					fmt.Println(startDay, endDay)
				}
			}
		}
	}
}
