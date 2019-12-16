package main

import (
	"fmt"
	"goLirary/db/mysql"
)

func main() {
	mysql := mysql.NewMysqlPool()
	fmt.Println(mysql)
}
