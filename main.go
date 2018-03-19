package main

import (
	"fmt"
	"goLirary/base/sysExample"
)

func main() {
	t := sysExample.NewOsDemo()
	fmt.Println(t.ReadEnvByKey("path"))
}
