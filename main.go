package main

import (
	"fmt"
	"go-api/db"
	"go-api/handler"
	"net/http"
)



func main(){

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS Profile (id integer not null primary key, Name text, Age integer);
	`
	db.CreateTable(sqlStmt)
	db.Db.Close()

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/users", handler.HandleUser)
	http_error := http.ListenAndServe("localhost:3003", nil)

	if http_error != nil{
		panic(http_error)
	}else{
		fmt.Println("Server is running on port: 3003")
	}
}