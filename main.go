package main

import (
	"fmt"
	"goLirary/base"
)

func main() {
	t := base.NewTimeDemo()
	fmt.Println(t.CurrentUnixTime())
}
