package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error: ", err)
	}

	a := Item{"first", "a first item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	err = client.Call("API.AddItem", a, &reply)
	err = client.Call("API.AddItem", b, &reply)
	err = client.Call("API.AddItem", c, &reply)
	err = client.Call("API.GetDB", "", &db)

	if err != nil {
		log.Fatal("rpc: ", err)
	}

	fmt.Println("Database: ", db)
}
