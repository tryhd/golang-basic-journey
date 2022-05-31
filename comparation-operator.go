package main

import "fmt"

func main() {
	var (
		name1 = "Fajar"
		name2 = "Dwi"
	)
	var result bool = name1 == name2

	var (
		value1 = 100
		value2 = 200
	)
	var result2 bool = value1 > value2

	fmt.Println(name1, "==", name2, "= ", result)
	fmt.Println(value1, ">", value2, "=", result2)
}
