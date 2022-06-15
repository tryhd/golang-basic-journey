package main

import "fmt"

//with param
func isGenap(nilai int) {
	if nilai%2 == 0 {
		fmt.Println(nilai, "adalah bilangan genap")
	} else {
		fmt.Println(nilai, "adalah bilangan ganjil")
	}
}

//non parameter
func sayHello() {
	fmt.Println("Hello, world!")
}

//return value
func checkName(name string) string {
	if name == "" {
		return "Nama tidak boleh kosong"
	}
	if len(name) < 3 {
		return "Nama terlalu pendek"
	}
	return "Nama sudah sesuai"
}

//return multiple value
func getBiodata() (string, string, int) {
	return "Oki", "Bandung", 25
}

func getBook() (string, string) {
	return "INI BUKU BARU", "BUDI"
}

//return with name return value
func getData() (name string, mk string, nilai float32) {
	name = "Oki Iskandar"
	mk = "Bahasa Indonesia"
	nilai = 3.0

	return
}

//variadic function
func getSumNilai(name string, nilai ...int) (string, int) {
	nilaiAkhir := 0

	for _, value := range nilai {
		nilaiAkhir += value
	}

	return name, nilaiAkhir
}

//fucntion value
func isValue(name string) string {
	return "Good Bye" + name
}

//function as parameter
type filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

func sayHelloWithFilter(name string, filter filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

//anonymous function
type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("User", name, "dilarang masuk")
	} else {
		fmt.Println("User", name, "boleh masuk")
	}
}

func main() {
	sayHello()
	print("\n")
	isGenap(10)
	isGenap(11)
	print("\n")
	result := checkName("oki")

	fmt.Println(result)
	name, address, age := getBiodata()

	fmt.Println("Nama :", name)
	fmt.Println("Address :", address)
	fmt.Println("Age :", age)

	title, _ := getBook()

	fmt.Println(title)

	mhs, mk, nilai := getData()

	fmt.Println("Nama :", mhs)
	fmt.Println("Mata Kuliah :", mk)
	fmt.Println("NIlai :", nilai)

	name, total := getSumNilai("Oki", 10, 10, 10, 10, 10)

	fmt.Println("Name :", name)
	fmt.Println("Nilai :", total)

	value := isValue

	fmt.Println(value("Oki"))

	sayHelloWithFilter("Oki", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	balaclist := func(name string) bool {
		return name == "admin"
	}

	registerUser("Oki", balaclist)
	registerUser("admin", balaclist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
