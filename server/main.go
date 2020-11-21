package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

var database []Item

type API int

func main() {
	var api = new(API)

	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving", err)
	}
}

func (a *API) GetDB(empty string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}

	*reply = getItem
	return nil
}

func (a *API) AddItem(i Item, reply *Item) error {
	fmt.Println(i)

	database = append(database, i)
	*reply = i
	return nil
}

func (a *API) EditItem(i Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == i.Title {
			database[idx] = Item{Title: i.Title, Body: i.Body}
			changed = database[idx]
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(i Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Title == i.Title && val.Body == i.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = i
			break
		}
	}

	*reply = del
	return nil
}
