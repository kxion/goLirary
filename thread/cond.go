package thread

import (
	"log"
	"sync"
	"time"
)

type CondDemo struct{}

func (s *CondDemo) Run() {
	con := sync.NewCond(&sync.Mutex{})

	cons := false

	go func() {
		con.L.Lock()
		time.Sleep(10 * time.Second)
		cons = true
		log.Printf("设置信号")
		con.Signal()
		con.L.Unlock()
	}()
	con.L.Lock()
	for !cons {
		log.Printf("等待信号")
		con.Wait()
	}
	con.L.Unlock()
	log.Printf("信号等待成功！")
}
