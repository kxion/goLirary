package process

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type SignalDemo struct{}

func NewSignalDemo() *SignalDemo {
	return new(SignalDemo)
}

/*
监听信号
*/
func (s *SignalDemo) Rundemo1() {
	sigRecv := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv, sigs...)
	for sig := range sigRecv {
		fmt.Printf("recived %s", sig)
		defer fmt.Println(123)
		os.Exit(0)
	}
}

//通过管道对于进行进程进行操作，这一块可以借助信号操作
func (s *SignalDemo) Rundemo2() {
	cmd1 := exec.Command("ps", "-aux")
	var output1 bytes.Buffer
	cmd1.Stdout = &output1
	cmd1.Start()
	cmd1.Wait()

	cmd2 := exec.Command("awk", "{print $2}")
	cmd2.Stdin = &output1
	var output2 bytes.Buffer
	cmd2.Stdout = &output2
	cmd2.Start()
	cmd2.Wait()
	temp := strings.Split(string(output2.Bytes()), "\n")
	for _, val := range temp[1:] {
		valInt, _ := strconv.Atoi(val)
		proc, err := os.FindProcess(valInt)
		fmt.Println(proc, err)
		//proc.Signal()
	}
}
