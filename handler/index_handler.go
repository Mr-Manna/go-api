package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

var message = []Message{
	{ Status: 200, Message: "Welcome to building apis with go!"},
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	text,errr := json.Marshal(message)
	if errr != nil{
		fmt.Print(errr)
	}
	w.Header().Add("Content-Type","application/json, charset=utf-8")
	w.Write(text)
}