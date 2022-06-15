package main

import "fmt"

func endApp() {
	fmt.Println("Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Memulai aplikasi")
}
func main() {
	fmt.Println("Error False")
	runApp(false)
	fmt.Println("Error True")
	runApp(true)
}
