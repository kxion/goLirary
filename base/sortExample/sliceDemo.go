package sortExample

import (
	"fmt"
	"sort"
)

type SliceDemo struct {
}

func NewSliceDemo() *SliceDemo {
	return &SliceDemo{}
}

type DemoData struct {
	Name     string
	Age      int64
	Money    float64
	AgeTrend int64
}

var demoInfo = []DemoData{{"Bob", 31, 20.04, 0.0}, {"jim", 10, 70.04, 0.0}, {"bim", 10, 200.04, 0.0}}
var compareInFO = []DemoData{{"Bob", 32, 23.04, 0.0}, {"jim", 15, 79.04, 0.0}, {"bim", 33, 100.04, 0.0}}

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

func (n *SliceDemo) SortDataMultiFiled(infoList []DemoData) []DemoData {
	sort.Slice(infoList, func(i, j int) bool {
		if infoList[i].Name > infoList[j].Name {
			return true
		} else {
			if infoList[i].Age > infoList[j].Age {
				return true
			}
		}
		return false
	})
	return infoList
}

func (n *SliceDemo) RunDemoOne() []DemoData {
	tempData := n.SortDataOneFiled(demoInfo)
	return tempData
}

func (n *SliceDemo) RunDemoMulti() []DemoData {
	tempData := n.SortDataMultiFiled(demoInfo)
	return tempData
}

func (n *SliceDemo) Run(data *[]DemoData) {
	for _, val := range *data {
		val.Age = val.Age
	}
	fmt.Println(demoInfo)
}

func (n *SliceDemo) GoRun() {
	for key, val := range demoInfo {
		tmp := compareInFO[key]
		val := n.ComputeRun(val, tmp)
		fmt.Println(val)
	}
}

func (n *SliceDemo) ComputeRun(originInfo, compareInfo DemoData) DemoData {
	originInfo.AgeTrend = originInfo.Age - compareInfo.Age
	return originInfo
}
