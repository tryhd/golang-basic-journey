package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Budi"
	names[1] = "Andi"
	names[2] = "Sri"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [5]int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(values)

	fmt.Println(len(values))
	fmt.Println(cap(values))
	fmt.Println(len(names))
	fmt.Println(cap(names))
}
