package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	fmt.Println("f = ", f)
	f += a

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
}
