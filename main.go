package main

import (
	"fmt"
	"goLirary/base/fileExample"
)

func main() {
	t := fileExample.NewReadFileDemo()
	info, _ := t.ReadBinaryFile()
	fmt.Println(info)
}
