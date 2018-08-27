package thread

import (
	"fmt"
	"sync"
)

/*
只执行一次
*/
type OnceDemo struct{}

func (s *OnceDemo) Run() {
	var count int
	var once sync.Once
	for i := 0; i < 4; i++ {
		once.Do(func() {
			count += 1
		})
	}
	fmt.Println(count)
}

func (s *OnceDemo) SingleDemo() *OnceDemo {
	var once sync.Once
	var ins *OnceDemo
	once.Do(func() {
		ins = &OnceDemo{}
	})
	return ins
}
