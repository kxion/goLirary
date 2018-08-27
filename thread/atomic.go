package thread

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type AtomicDemo struct{}

func (s *AtomicDemo) NoLockAddRun() {
	var total int64
	for i := 0; i < 10; i++ {
		go func() {
			total += 1
			fmt.Println(total)
		}()
	}
}

func (s *AtomicDemo) LockAddRun() {
	var total int64
	var nx sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			nx.Lock()
			total += 1
			fmt.Println(total)
			nx.Unlock()
		}()
	}
}

func (s *AtomicDemo) AutomicAddRun() {
	var total int64
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&total, 1)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("The total number is", total)
}

/*
	原子值是值传递,不是地址传递
*/
func (s *AtomicDemo) RunAtomicValue() {
	var countV atomic.Value
	countV.Store([]int{1, 2, 3, 4})
	s.StoreAtomicValue(&countV)
	fmt.Println(countV.Load())
}

func (s AtomicDemo) StoreAtomicValue(conv *atomic.Value) {
	conv.Store([]int{6, 7, 8, 9})
}
