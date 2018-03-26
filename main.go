package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func f2(in chan int) {
	in <- 2
}

func main() {
	out := make(chan int)
	go f1(out)
	go f2(out)
	fmt.Println(<-out)
}
