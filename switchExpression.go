package main

import "fmt"

func main() {
	name := "Oki"

	switch name {
	case "Budi":
		fmt.Println("Hello", name, "!")
	case "Andi":
		fmt.Println("Hello Andi!")
	default:
		fmt.Println("Hello, stranger!")
	}

	switch length := len(name); length < 3 {
	case true:
		fmt.Println("Your name is", length, "characters short. Incorrect!")
	default:
		fmt.Println("Your name is", length, "characters long. Correct!")
	}

	length := len(name)

	switch {
	case length < 3:
		fmt.Println("Nama Terlalu Pendek!")
	case length > 3:
		fmt.Println("Nama Sama Dengan 3 Karakter!")
	default:
		fmt.Println("Nama Sama Dengan 3 Karakter!")

	}
}
