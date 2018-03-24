package fileExample

import (
	"bufio"
	"fmt"
	"os"
)

type ReadFileDemo struct{}

func NewReadFileDemo() *ReadFileDemo {
	return &ReadFileDemo{}
}

func (r *ReadFileDemo) ReadFileLine() []string {
	fileObj, readErr := os.Open("data/file/readDemo.log")
	if readErr != nil {
		fmt.Println(readErr)
	}
	defer fileObj.Close()
	var infoList []string

	//这一块因为*File实现 io.Reader的方法，所以这一块是可以调用
	fileInfos := bufio.NewScanner(fileObj)
	for fileInfos.Scan() {
		infoList = append(infoList, fileInfos.Text())
	}
	return infoList
}

func (r *ReadFileDemo) ReadBinaryFile() ([]byte, error) {
	fileObj, readErr := os.Open("data/file/test.bin")
	if readErr != nil {
		fmt.Println(readErr)
	}
	defer fileObj.Close()

	//获取二进制文件的字节长度
	stats, err := fileObj.Stat()
	var size int64 = stats.Size()

	//计算文件的字节长度
	infoList := make([]byte, size)
	if err != nil {
		return infoList, err
	}

	bufr := bufio.NewReader(fileObj)
	_, err = bufr.Read(infoList)
	if err != nil {
		return infoList, err
	}
	return infoList, nil
}
