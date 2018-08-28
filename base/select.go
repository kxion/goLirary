package base

import (
	"fmt"
	"time"
)

type SelectDemo struct{}

func (s *SelectDemo) Run() {
	ch1 := make(chan int, 0)
	ch2 := make(chan byte, 0)
	go s.SendChan1(ch1)
	go s.SendChan2(ch2)
	go s.RunDemo(ch1, ch2)
}

func (s *SelectDemo) RunDemo(ch1 chan int, ch2 chan byte) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("output chan1 %d\n", v)
		case v := <-ch2:
			s := []byte{v}
			fmt.Printf("output chan2 %s\n", string(s))
		default:
			time.Sleep(10 * time.Second)
			fmt.Println("output no value")
		}
	}
}

func (s *SelectDemo) SendChan1(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
}

func (s *SelectDemo) SendChan2(ch chan byte) {
	s1 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'b', '3', '2', '4', '5', '6', '7', '8', '9', '0', 'k'}
	for _, v := range s1 {
		ch <- v
	}
}

func (s *SelectDemo) SendRun() {
	ch1 := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case ch1 <- i:
				fmt.Println("input %d", i)
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("output chan1 %d", <-ch1)
		}
	}()
}
