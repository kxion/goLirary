package library

import (
	"bytes"
	"fmt"
)

type ByteDemo struct{}

func (s *ByteDemo) Run() {
	var data bytes.Buffer
	data.WriteString("123")
	data.WriteByte('\n')
	fmt.Fprintf(&data, "%s", "123")
	fmt.Println(string(data.Bytes()))
}
