package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Memulai aplikasi")
	result := value / 10
	fmt.Println("Hasil :", result)
}
func main() {
	runApplication(5)
}
