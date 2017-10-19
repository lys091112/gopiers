package database

import _ "github.com/go-sql-driver/mysql"
import "database/sql"
import "fmt"

func GetConn() (*sql.DB, error) {
	return sql.Open("mysql", "root:12345678@tcp(localhost:13306)/test?charset=utf8")
}

func Insert(db *sql.DB, event alertEvent) {
	stat, err := db.Prepare("insert into persion (name) values  (?)")
	checkErr(err)

	res, err := stat.Exec("xx")
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(num)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
