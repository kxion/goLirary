package main

import (
	"fmt"
	"goLirary/base/timeExample"
	"time"
)

func main() {
	t := timeExample.NewTimeDemo()
	day := "2017-10-15"
	temp, err := time.Parse("2006-01-02", day)
	fmt.Println(temp.Weekday().String(), err)
	fmt.Println(t.GetWeek(temp))
}
