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
	var infoList []byte
	_, err := fileObj.Read(infoList)
	if err != nil {
		return infoList, err
	}
	return infoList, nil
}
