package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Oki")
	data.PushBack("Oki 1")
	data.PushBack("Oki 2")
	//depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
