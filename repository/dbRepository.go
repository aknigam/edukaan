package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

/*
	connect to the Db
	sql.Open does not open a connection. A connection pool is maintained
*/
func init() {
	var err error
	Db, err = sql.Open("mysql", "edukan:edukan@/edukaan")
	if err != nil {
		panic(err)
	}
}
