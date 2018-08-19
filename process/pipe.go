package process

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Pipe struct{}

func NewPipe() *Pipe {
	return new(Pipe)
}

func (s *Pipe) RunCmd(parameter1, parameter2 string) {
	cmd := exec.Command(parameter1, parameter2) //获取cmd对象
	stdout, err := cmd.StdoutPipe()             //获取输出管道
	if err != nil {
		log.Printf("output err:%s", err.Error())
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("error command:%s", err.Error())
		return
	}
	outputBuf0 := bufio.NewReader(stdout)
	for {
		outputo, err := outputBuf0.ReadString('\n')
		fmt.Printf("%s", outputo)
		//读取完毕
		if err == io.EOF {
			break
		}
		//读取失败
		if err != nil {
			fmt.Printf("read file failed, err:%v", err)
			break
		}
	}
}

//类似于执行ps aux | grep hellow
func (s *Pipe) RunGrepCmd(cmd1Parameter1, cmd1Parameter2, cmd2Parameter1, cmd2Parameter2 string) {
	cmd1 := exec.Command(cmd1Parameter1, cmd1Parameter2)
	var output1 bytes.Buffer
	cmd1.Stdout = &output1
	cmd1.Start()
	cmd1.Wait()
	//以cmd1的输出做为cmd2的输出
	cmd2 := exec.Command(cmd2Parameter1, cmd2Parameter2)
	cmd2.Stdin = &output1
	var output2 bytes.Buffer
	cmd2.Stdout = &output2
	cmd2.Start()
	cmd2.Wait()
	fmt.Printf("%s\n", string(output2.Bytes()))
}

/*
命名管道必须要成对出现，不然会被阻塞,单独的reader和writer都会阻塞起来，
os.pipe的并发的时候，需要加锁
cmd := process.NewPipe()
	reader, writer, _ := os.Pipe()
	go func() {
		cmd.WritePipe(writer)
	}()
	go func() {
		cmd.ReadPipe(reader)
	}()
time.Sleep(10 * time.Second)
*/
//命名管道
func (s *Pipe) WritePipe(f *os.File) {
	var input = []byte{'1', '2', '3'}
	f.Write(input)
}

func (s *Pipe) ReadPipe(f *os.File) {
	output := make([]byte, 10)
	f.Read(output)
	fmt.Printf("output:%s", string(output))
}

/*
	io的pipe模块
*/

func (s *Pipe) WriteIoPipe(f *io.PipeWriter) {
	f.Write([]byte{'1', '2', '3'})
}

func (s *Pipe) ReadIoPipe(f *io.PipeReader) {
	output := make([]byte, 10)
	f.Read(output)
	fmt.Printf("output:%s", string(output))
}
