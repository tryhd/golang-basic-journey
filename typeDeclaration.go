package main

import "fmt"

func main() {

	type noKtp string
	type married bool

	var nik noKtp = "320831199502022001"
	var isMarried married = true

	fmt.Println("nik = ", nik)
	fmt.Println("isMarried = ", isMarried)
}
