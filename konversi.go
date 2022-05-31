package main

import "fmt"

func main() {
	var (
		nilai32 = 129
		nilai64 = int64(nilai32)
		nilai8  = int8(nilai32)
	)

	var (
		name           = "Fajar"
		f       byte   = name[0]
		fString string = string(f)
	)

	fmt.Println("nilai32 = ", nilai32)
	fmt.Println("nilai64 = ", nilai64)
	fmt.Println("nilai8 = ", nilai8)

	fmt.Println(name)
	fmt.Println("f = ", f)
	fmt.Println("fString = ", fString)
}
