package database

import _ "github.com/go-sql-driver/mysql"
import "database/sql"

func GetConn() (*sql.DB, error) {
	return sql.Open("mysql", "root:12345678@tcp(localhost:13306)/alert_engine?charset=utf8")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
