package main

import (
	"fmt"
	"goLirary/base/fileExample"
)

func main() {
	t := fileExample.NewWriteFileDemo()
	var textInfo []string
	for i := 0; i < 50; i++ {
		val := fmt.Sprintf("%d", i)
		textInfo = append(textInfo, val)
	}
	err := t.WriteTextFile("data/file/write.txt", textInfo)
	fmt.Println(err)
}
