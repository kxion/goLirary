package structExample

type StructDemo struct{}

type DemoStruct struct {
	Uid     int64
	Name    string
	Age     float64
	Sex     string
	Address string
}

func NewStructDemo() *StructDemo {
	return &StructDemo{}
}

func (s *StructDemo) Run() {

}
