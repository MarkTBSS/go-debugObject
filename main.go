package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func debugObj(obj interface{}) {
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(raw))
}

func main() {
	bookInstance := Book{
		Id:     1,
		Name:   "bookWay1",
		Author: "bookAuthor1",
	}
	fmt.Println(bookInstance)
	debugObj(bookInstance)
}
