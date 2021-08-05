package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{ Name: "Chandan Manna", Age: 34, Email: "chandan.manna@gmail.com"},
	{ Name: "Koushik Manna", Age: 30, Email: "koushik.manna@gmail.com"},
}

func HandleUser(w http.ResponseWriter, r *http.Request){
	users, errr := json.Marshal(users)
	if errr != nil{
		fmt.Print(errr)
	}
	w.Header().Add("Content-Type","application/json, charset=utf-8")
	w.Write(users)
}
