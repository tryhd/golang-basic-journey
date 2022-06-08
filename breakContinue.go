package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Ini perulangan ke", i)
	}

	for i := 11; i <= 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Ini perulangan ke", i)
	}
}
