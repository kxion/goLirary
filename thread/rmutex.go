package thread

import (
	"fmt"
	"sync"
)

type RMuTexDemo struct{}

var test = make(map[int]int)

var rw *sync.RWMutex

func init() {
	rw = new(sync.RWMutex)
}

func (s *RMuTexDemo) Run() {
	go s.Add()
	go s.Read()
}

func (s *RMuTexDemo) Add() {
	rw.Lock()
	for i := 0; i < 20000; i++ {
		test[i] = i
	}
	defer rw.Unlock()
}

func (s *RMuTexDemo) Read() {
	rw.RLock()
	for i := 0; i < 20000; i++ {
		if tmp, ok := test[i]; ok {
			fmt.Println(tmp)
		}
	}
	defer rw.RUnlock()
}
