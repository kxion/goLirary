package process

import (
	"fmt"
	"sync"
)

type Semaphore struct{}

func NewSemaphore() *Semaphore {
	return new(Semaphore)
}

/*
保证每个协程都执行完毕
*/
func (s *Semaphore) RunSemaphore() {
	var wg sync.WaitGroup
	wg.Add(2)
	var n, m int64
	go func() {
		n += 100
		wg.Done()
	}()
	go func() {
		m += 200
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(m, n)
}
