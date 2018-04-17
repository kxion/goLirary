package structExample

import "fmt"

type StructDemo struct{}

type DemoStruct struct {
	Uid     int64
	Name    string
	Age     int64
	Sex     string
	Address string
}

func NewStructDemo() *StructDemo {
	return &StructDemo{}
}

//如果需要改变结构体的数值，这一块需要参考如下图的用法，使用指针去操作
func (s *StructDemo) ChangeStructValue() {
	var data = []*DemoStruct{}
	for i := 1; i < 10; i++ {
		var infoData = DemoStruct{Age: int64(i), Uid: int64(i)}
		data = append(data, &infoData)
	}

	for _, val := range data {
		fmt.Println(val.Uid, val.Age)
	}
	for _, val := range data {
		val.Age = val.Age + 1
		val.Uid = val.Uid + 1
	}
	fmt.Println("*******************")
	for _, val := range data {
		fmt.Println(val.Uid, val.Age)
	}
}

func (s *StructDemo) Run() {
	s.ChangeStructValue()
}
