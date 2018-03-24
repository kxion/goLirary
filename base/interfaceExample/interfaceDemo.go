package interfaceExample

import "fmt"

type InterfaceDemo struct{}

func NewInterfaceDemo() *InterfaceDemo {
	return &InterfaceDemo{}
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c = *c + ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (i *InterfaceDemo) Run() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	name := "nihao"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
