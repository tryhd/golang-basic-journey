package main

import (
	"fmt"
)

func main() {

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[1:4]

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)

	months[2] = "Diganti"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(len(slice2))

	var slice3 = append(slice2, "Tambahan")

	fmt.Println(slice3)

	slice3[1] = "Ganti Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "A"
	newSlice[1] = "B"
	fmt.Println(newSlice)

	fromSlice := make([]string, len(newSlice), cap(newSlice))

	copy(fromSlice, newSlice)
	fmt.Println(fromSlice)

	//perbedaan array dan slice
	isArray := [5]int{1, 2, 3, 4, 5}
	isSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(isArray)
	fmt.Println(isSlice)
}
