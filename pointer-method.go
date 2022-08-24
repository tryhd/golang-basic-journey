package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	oki := Man{"Oki"}
	oki.Married()

	fmt.Println(oki.Name)
}
