package main

import "fmt"

type Customer struct {
	Name    string
	Age     int
	Address string
}

//struct function
func (c Customer) sayHello() {
	fmt.Println("Hello, my name is", c.Name)
}

func main() {
	c := Customer{
		Name:    "John",
		Age:     30,
		Address: "Jl. Kebon Kacang",
	}
	fmt.Println(c)

	newCust := Customer{"Joko", 70, "Jl. Kebon Kacang II"}
	fmt.Println(newCust)

	c.sayHello()
}
