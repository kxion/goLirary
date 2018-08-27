package thread

import (
	"fmt"
	"sync"
)

type MutexDemo struct{}

func (s *MutexDemo) Run() {
	var m sync.Mutex
	for i := 1; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i)
			m.Unlock()
		}(i)
	}
}
