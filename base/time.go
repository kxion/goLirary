package base

import (
	"fmt"
	"time"
)

type TimeDemo struct{}

func (s *TimeDemo) RunTimer1() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("present time:%s\n", time.Now().Format("2006-01-02 15:04:05"))
	expireTime := <-timer.C
	fmt.Printf("%s", expireTime.Format("2006-01-02 15:04:05"))
	timer.Stop()
}

//定义超时的操作
func (s *TimeDemo) RunTimer2() {
	iniChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second)
			iniChan <- i
		}
		close(iniChan)
	}()
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(time.Second)
		} else {
			timer.Reset(time.Second)
		}
		select {
		case v, ok := <-iniChan:
			if !ok {
				fmt.Println("end")
				return
			}
			fmt.Println(v)
		case <-timer.C:
			fmt.Println("timeOut")
		}
	}
}

func (s *TimeDemo) RunTicker1() {
	iniChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			select {
			case iniChan <- 1:
			case iniChan <- 2:
			case iniChan <- 3:
			}
		}
		fmt.Println("send end")
	}()
	go func() {
		var sem int
		for v := range iniChan {
			sem += v
			fmt.Println(v, sem)
			if sem > 20 {
				ticker.Stop()
			}
		}
		fmt.Println("recv end")
	}()
}
