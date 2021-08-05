package main

import (
	"fmt"
	"go-api/db"
	"go-api/handler"
	"net/http"
)



func main(){

	sqlStmt := `
		create table Profile (id integer not null primary key, Name text, Age integer);
	`
	db.Db.Exec(sqlStmt)

	// if err != nil {
	// 	fmt.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }

	http.HandleFunc("/", handler.IndexHandler)
	http_error := http.ListenAndServe("localhost:3003", nil)

	if http_error != nil{
		panic(http_error)
	}else{
		fmt.Println("Server is running on port: 3003")
	}
}