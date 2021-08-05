package db

import (
	"database/sql"
	// "fmt"
	_ "github.com/mattn/go-sqlite3"
)


var Db, _ = sql.Open("sqlite3", "./main.db")

// func(*db_error ) (error string) {

// 	if &db_error != nil {
// 		panic(db_error)
// 	}
// }()
