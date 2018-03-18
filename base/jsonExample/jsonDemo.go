package jsonExample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonDemo struct{}

func NewJsonDemo() *JsonDemo {
	return &JsonDemo{}
}

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func (*JsonDemo) EncodeDemo1() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	paJson, _ := json.Marshal(pa)
	v1 := make(map[int]Address)
	for i := 0; i < 50; i++ {
		v1[i] = *pa
	}
	v1Js, _ := json.Marshal(v1)
	fmt.Printf("%s", v1Js)
	fmt.Printf("pa %s", paJson)
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
}

func (*JsonDemo) ReadJsonFile() {
	fileObj, _ := ioutil.ReadFile("data/json/demo1.json")
	var decodeJs []Address
	err := json.Unmarshal(fileObj, &decodeJs)
	for key, val := range decodeJs {
		fmt.Println(key, val.City)
	}
	fmt.Println(err)

}
