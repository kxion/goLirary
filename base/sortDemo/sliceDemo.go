package sortDemo

import (
	"sort"
)

type SliceDemo struct {
}

func NewSliceDemo() *SliceDemo {
	return &SliceDemo{}
}

type DemoData struct {
	Name  string
	Age   int64
	Money float64
}

var demoInfo = []DemoData{{"Bob", 31, 20.04}, {"jim", 10, 70.04}}

//这一块主要是按照三种类型进行演示
func (n *SliceDemo) SortDataOneFiled(infoList []DemoData) []DemoData {
	sort.Slice(infoList, func(i, j int) bool {
		if infoList[i].Name > infoList[i].Name {
			return true
		}
		return false
	})
	return infoList
}

func (n *SliceDemo) RunDemo() []DemoData {
	tempData := n.SortDataOneFiled(demoInfo)
	return tempData
}
