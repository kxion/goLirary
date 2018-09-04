package interfaceExample

import (
	"fmt"
	"reflect"
)

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

type Test1 struct {
	Age int
}

type Test2 struct {
	Age int
}

func (i *InterfaceDemo) Run2() {
	data1 := []Test1{{Age: 10}, {Age: 10}, {Age: 10}, {Age: 10}}
	fmt.Println(i.GetLen(data1))
	data2 := []int{1, 2, 3}
	fmt.Println(i.GetLen(data2))
	data3 := []string{"1", "2"}
	fmt.Println(i.GetLen(data3))
	data4 := []Test2{{Age: 10}, {Age: 10}, {Age: 10}, {Age: 10}, {Age: 10},
		{Age: 10}, {Age: 10}, {Age: 10}}
	fmt.Println(i.GetLen(data4))
}

func (i *InterfaceDemo) GetLen(data interface{}) int {
	var sliceLen int
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(data)
		sliceLen = s.Len()
	}
	return sliceLen
}
