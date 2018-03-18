package base

type MapDemo struct{}

func NewMapDemo() *MapDemo {
	return &MapDemo{}
}

type MapDemo1 struct {
	Name  string
	Age   int64
	Id    int64
	Money float64
}

var DemoInfo = []MapDemo1{{"Bob", 31, 1, 20.04}, {"jim", 10, 10, 70.04}, {"bim", 10, 30, 200.04}}

//slice转换成标准的key value的结构
//类似与php的array["1"=>[1,2,3]]
func (m *MapDemo) ObJectValue() map[int64]MapDemo1 {
	var obj = make(map[int64]MapDemo1)
	for _, val := range DemoInfo {
		obj[val.Id] = val
	}
	return obj
}

//按照数组下标去生成的数组
func (m *MapDemo) NormalValue() map[int]MapDemo1 {
	var obj = make(map[int]MapDemo1)
	for index, val := range DemoInfo {
		obj[index] = val
	}
	return obj
}
