package helper

import "fmt"

var version = 1.0
var Application = "Belajar GoLang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("Hello ", name, " Good Bye")
}
