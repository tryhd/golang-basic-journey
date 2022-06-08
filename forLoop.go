package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Ini perulangan ke", counter)
		counter++
	}

	for i := 11; i <= 20; i++ {
		fmt.Println("Ini perulangan ke", i)
	}

	slice := []string{"Budi", "Andi", "Sri"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println(index, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Budi"
	person["age"] = "20"
	person["address"] = "Jakarta"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
