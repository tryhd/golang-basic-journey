package main

import "fmt"

func main() {
	name := "Oki"

	if name == "Budi" {
		fmt.Println("Hello", name, "!")
	} else if name == "Andi" {
		fmt.Println("Hello Andi!")
	} else {
		fmt.Println("Hello, stranger!")
	}

	if length := len(name); length > 3 {
		fmt.Println("Your name is", length, "characters long.")
	} else {
		fmt.Println("Your name is", length, "characters short.")
	}
}
