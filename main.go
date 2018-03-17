package main

import (
	"fmt"
	"goLirary/base/sortDemo"
)

func main() {
	t := sortDemo.NewSliceDemo()
	fmt.Println(t.RunDemoOne())
	fmt.Println(t.RunDemoMulti())
}
