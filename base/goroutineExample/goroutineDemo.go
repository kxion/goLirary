package goroutineExample

import (
	"fmt"
)

type GoRoutineDemo struct{}

func NewGoRoutineDemo() *GoRoutineDemo {
	return &GoRoutineDemo{}
}

func (g *GoRoutineDemo) PrintH() {
	fmt.Println("H")
}

func (g *GoRoutineDemo) PrintW() {
	fmt.Println("W")
}
