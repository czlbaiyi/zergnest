package mysqldb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init(user string, pwd string) {
	dbstr := user + ":" + pwd + "@/database"
	var err error
	db, err = sql.Open("mysql", dbstr)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
}

func close() {
	db.Close()
}
