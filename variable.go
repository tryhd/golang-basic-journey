package main

import "fmt"

func main() {
	var name string
	var age = 20
	hobby := "baca buku"
	name = "Budi"

	var (
		country = "Indonesia"
		city    = "Jakarta"
	)

	fmt.Println("Hello,", name)
	fmt.Println("Your age is", age)
	fmt.Println("Your hobby is", hobby)
	fmt.Println("You live in", country, "in", city)
}
