package fileExample

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
	"runtime"
)

type payload struct {
	One   float32
	Two   float64
	Three uint32
}

type WriteFileDemo struct{}

func NewWriteFileDemo() *WriteFileDemo {
	return &WriteFileDemo{}
}

func (w *WriteFileDemo) WriteByte(file *os.File, binInfo []byte) error {
	_, err := file.Write(binInfo)
	return err
}

func (w *WriteFileDemo) WriteText(file *os.File, text string) error {
	_, err := file.WriteString(text)
	return err
}

func (w *WriteFileDemo) WriteTextFile(dir string, content []string) error {
	createTextFile, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer createTextFile.Close()

	//获取当前运行机器的操作系统
	var newLineFlag string
	if runtime.GOOS == "windows" {
		newLineFlag = "\r\n"
	} else {
		newLineFlag = "\n"
	}
	outputWriter := bufio.NewWriter(createTextFile)
	for _, val := range content {
		outputWriter.WriteString(val + newLineFlag)
	}
	outputWriter.Flush()
	return nil
}

func (w *WriteFileDemo) WriteBinaryFile(dir string) error {
	createBinFile, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer createBinFile.Close()
	for i := 0; i < 50; i++ {
		s := &payload{
			1.10,
			2.20,
			3,
		}
		var binBuf bytes.Buffer
		binary.Write(&binBuf, binary.BigEndian, s)
		w.WriteByte(createBinFile, binBuf.Bytes())
	}
	return nil
}
