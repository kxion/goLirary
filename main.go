package main

import (
	"fmt"
	"goLirary/base/goroutineExample"
	"time"
)

func main() {
	fmt.Println("start")
	t := goroutineExample.NewChannelDemo()
	var ch = make(chan string)
	go t.SendData(ch)
	go t.RecieveData(ch)
	time.Sleep(10 * 1e9)
	fmt.Println("end")
}
