package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)


var Db, _ = sql.Open("sqlite3", "./main.db")

// func(*db_error ) (error string) {

// 	if &db_error != nil {
// 		panic(db_error)
// 	}
// }()
 func CreateTable(statement string){

	_, error := Db.Exec(statement)

	if error != nil {
		fmt.Printf("%q: %s\n", error, statement)
	}else{
		fmt.Println( "Table created successfully.")
	}
 } 