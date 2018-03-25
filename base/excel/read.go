package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/*
excel的扩展采用的是https://github.com/360EntSecGroup-Skylar/excelize
*/

type Read struct{}

func NewRead() *Read {
	return &Read{}
}

func (r *Read) ReadExcel(path string, sheetName string) {
	xlsx, _ := excelize.OpenFile(path)
	rows := xlsx.GetRows(sheetName)
	//var listInfo []User
	for _, row := range rows {
		//var user User
		for _, colCell := range row {
			fmt.Println(colCell)
		}
		//listInfo = append(listInfo, user)
	}

}
