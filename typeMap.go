package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Budi",
		"age":     "20",
		"address": "Jakarta",
	}
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "Budi"
	book["publisher"] = "Gramedia"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "author")
	fmt.Println(book)
	fmt.Println(len(book))
}
