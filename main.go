package main

import (
	"goLirary/base"
	"time"
)

func main() {
	var s base.TimeDemo
	s.RunTicker1()
	time.Sleep(20 * time.Second)
}
