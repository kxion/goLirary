package excel

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type User struct {
	Name    string
	Age     int32
	Address string
	Email   string
	Like    string
}

var demoUserInfo = []User{
	{"zhang", 13, "杭州", "1101@qq.com", "台球"},
	{"zhang1", 14, "郑州", "1102@qq.com", "足球"},
	{"zhang2", 15, "苏州", "1103@qq.com", "蓝球"},
}
var demoColumnInfo = []string{"姓名", "年龄", "地址", "邮箱", "喜好"}

/*NewWrite
excel的扩展采用的是https://github.com/360EntSecGroup-Skylar/excelize
*/

type Write struct{}

func NewWrite() *Write {
	return &Write{}
}

func (w *Write) ConvertColumn(startColumn rune, columnInfo []string) map[string]string {
	var excelColumn = make(map[string]string)
	writeLen := len(columnInfo)
	tempAscii := int(startColumn)
	for i := 0; i < writeLen; i++ {
		tempKey := string(rune(tempAscii)) + "1"
		excelColumn[tempKey] = columnInfo[i]
		tempAscii = tempAscii + 1
	}
	return excelColumn
}

func (w *Write) ConVertRows(startColumn rune, rows []User, start string) map[string]map[string]interface{} {
	var nameInfo = make(map[string]interface{})
	var ageInfo = make(map[string]interface{})
	var addressInfo = make(map[string]interface{})
	var emailInfo = make(map[string]interface{})
	var likeInfo = make(map[string]interface{})
	writeLen := len(rows)

	var rowsData = make(map[string]map[string]interface{})
	for i := 0; i < writeLen; i++ {

		tempAscii := int(startColumn)
		tempUserInfo := rows[i]
		nameKey := string(rune(tempAscii)) + start
		nameInfo[nameKey] = tempUserInfo.Name
		tempAscii = tempAscii + 1
		ageKey := string(rune(tempAscii)) + start
		ageInfo[ageKey] = tempUserInfo.Age
		tempAscii = tempAscii + 1
		addressKey := string(rune(tempAscii)) + start
		addressInfo[addressKey] = tempUserInfo.Address
		tempAscii = tempAscii + 1
		emailKey := string(rune(tempAscii)) + start
		emailInfo[emailKey] = tempUserInfo.Email
		tempAscii = tempAscii + 1
		likeKey := string(rune(tempAscii)) + start
		likeInfo[likeKey] = tempUserInfo.Like
		tempLen, _ := strconv.Atoi(start)
		start = strconv.Itoa(tempLen + 1)
	}

	rowsData["name"] = nameInfo
	rowsData["age"] = ageInfo
	rowsData["address"] = addressInfo
	rowsData["email"] = emailInfo
	rowsData["like"] = likeInfo

	return rowsData
}

func (w *Write) Run() {
	rowDatas := w.ConVertRows('A', demoUserInfo, "2")
	columnDatas := w.ConvertColumn('A', demoColumnInfo)
	xlsx := excelize.NewFile()
	xlsx.NewSheet("用户信息")
	for k, v := range columnDatas {
		fmt.Println(k, v)
		xlsx.SetCellValue("用户信息", k, v)
	}
	for _, val := range rowDatas {
		for k, v := range val {
			fmt.Println(k, v)
			xlsx.SetCellValue("用户信息", k, v)
		}
	}
	xlsx.DeleteSheet("Sheet1")
	xlsx.SaveAs("./写入练习.xlsx")
}
