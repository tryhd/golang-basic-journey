package main

import "fmt"

type HasName interface {
	getName() string
}

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Name string
}

func sayHelloInterface(h HasName) {
	fmt.Println("Hello, my name is", h.getName())
}

func (animal Animal) getName() string {
	return animal.Name
}

func (person Person) getName() string {
	return person.Name
}

func main() {
	p := Person{"John", 30}
	sayHelloInterface(p)

	a := Animal{"Dog"}
	sayHelloInterface(a)
}
