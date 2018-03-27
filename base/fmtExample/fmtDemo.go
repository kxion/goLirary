package fmtExample

import (
	"bytes"
	"fmt"
)

type FmtDemo struct{}

func NewFmtDemo() *FmtDemo {
	return &FmtDemo{}
}

func (f *FmtDemo) Run() {
	var data bytes.Buffer
	fmt.Fprintf(&data, "%s", "hellow") //字符串转[]bytes
	fmt.Println(string(data.Bytes()))
}
