package convert

import (
	"fmt"
	"strconv"
)

type StringToNumber struct{}

func NewStringToNumber() *StringToNumber {
	return new(StringToNumber)
}

func (s *StringToNumber) StrInt() {
	str := "12"
	//convert to int
	intVal, _ := strconv.Atoi(str)
	fmt.Println(intVal)

	//convert to int64
	int64Val, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(int64Val)
	var intInfo int = 2
	var int64Info int64 = 4

	strVal := strconv.Itoa(intInfo)
	str64Val := strconv.FormatInt(int64Info, 10)
	fmt.Println(strVal, str64Val)
}

//科学计数法的抓换
func (s *StringToNumber) ComputeSc() {
	var computeSc float64 = 1.000000002e+09
	var floatInfo string
	floatInfo = strconv.FormatFloat(computeSc, 'f', 2, 64)
	fmt.Println(floatInfo)
}
