package thread

import (
	"fmt"
	"sync"
)

type WaitGroupDemo struct{}

func (s *WaitGroupDemo) Run() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		wg.Done()
		fmt.Printf("go runtime %s\n", "1")
	}()

	go func() {
		wg.Done()
		fmt.Printf("go runtime %s\n", "2")
	}()

	go func() {
		wg.Done()
		fmt.Printf("go runtime %s\n", "2")
	}()
	wg.Wait()
	fmt.Println("WaitGroupDemo finish demo")
}
