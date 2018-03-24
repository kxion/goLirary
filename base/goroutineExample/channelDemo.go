package goroutineExample

import "fmt"

type ChannelDemo struct{}

func NewChannelDemo() *ChannelDemo {
	return &ChannelDemo{}
}

func (c *ChannelDemo) SendData(ch chan string) {
	fmt.Println("begin send")
	ch <- "send begin"
	ch <- "hello"
	ch <- "fine"
	ch <- "end connect"
}

func (c *ChannelDemo) RecieveData(ch chan string) {
	fmt.Println("start recieve")
	for {
		input := <-ch
		fmt.Println(input)
	}
}
