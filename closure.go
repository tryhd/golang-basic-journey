package main

import "fmt"

func main() {
	counter := 0
	name := "Oki Iskandar"
	increment := func() {
		name := "Oki Iskandar Ganti"
		fmt.Println("Increment :", counter)
		fmt.Println("Name :", name)
		counter++
	}

	increment()
	increment()
	fmt.Println(name)
}
