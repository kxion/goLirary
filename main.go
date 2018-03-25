package main

import (
	"goLirary/base/excel"
)

func main() {
	t := excel.NewRead()
	t.ReadExcel("./写入练习.xlsx", "用户信息")
}
