package mapExample

import (
	"fmt"
)

type MapDemo struct{}

func NewMapDemo() *MapDemo {
	return &MapDemo{}
}

type MapDemo1 struct {
	Name  string
	Age   int64
	Id    int64
	Money float64
	Type  string
}

type MapDemo2 struct {
	Name  string
	Info  map[string]int64
	Total int64
}

var DemoInfo = []MapDemo1{{"Bob", 31, 1, 20.04, "type1"}, {"jim", 10, 10, 70.04, "type1"}, {"bim", 10, 30, 200.04, "type2"}}

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

func (m *MapDemo) UnionStruct() []MapDemo2 {
	var infos []MapDemo2
	var data = make(map[string]MapDemo2)
	for _, val := range DemoInfo {
		info, ok := data[val.Type]
		if !ok {
			info.Info = make(map[string]int64)
		}
		info.Name = val.Type
		info.Info[val.Name] = val.Age
		data[info.Name] = info
	}
	fmt.Println(data)
	return infos
}

//输出map结构
func (m *MapDemo) KeyValue() {
	var datas []map[string]int64
	for _, val := range DemoInfo {
		var info = make(map[string]int64)
		info[val.Name] = val.Age
		datas = append(datas, info)
	}
	fmt.Println(datas)
}

//输出标准的结构map[key][]struct
/*
var info = make(map[string][]struct)
info[key1]= append(info[key],val)
*/
