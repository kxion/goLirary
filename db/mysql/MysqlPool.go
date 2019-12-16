package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var goLibrayDB *sql.DB

type MysqlPool struct {
}

func NewMysqlPool() *MysqlPool {
	return new(MysqlPool)
}

func init() {
	mysqlPool := NewMysqlPool()
	mysqlPool.SetGoLibray()
}

func (s *MysqlPool) SetGoLibray() {
	goLibrayDB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golibary?charset=utf8&parseTime=true")
	if err != nil {
		log.Printf("数据库连接出错: %s", err.Error())
		return
	}
	goLibrayDB.SetMaxOpenConns(10)
	goLibrayDB.SetMaxIdleConns(5)
	goLibrayDB.SetConnMaxLifetime(28000 * time.Second)
}

func (s *MysqlPool) GetGolibray() *sql.DB {
	return goLibrayDB
}
