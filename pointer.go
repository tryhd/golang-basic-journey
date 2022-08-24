package main

import "fmt"

type Address struct {
	City    string
	Provice string
	Country string
}

//pointer function
func changeConuntyToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

//pointer method

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}

	var address3 = new(Address)
	address3.City = "Jakarta"
	address3.Provice = "DKI Jakarta"
	address3.Country = "Indonesia"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var alamat = Address{
		City:    "Bsndung",
		Provice: "Jawa Barat",
		Country: "",
	}
	changeConuntyToIndonesia(&alamat)
	fmt.Println(alamat)
}
