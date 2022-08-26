package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`o([a-z])i`)

	fmt.Println(regex.MatchString("oki"))
	fmt.Println(regex.MatchString("oci"))
	fmt.Println(regex.MatchString("oli"))

	fmt.Println(regex.FindAllString("oli oki ogi oka oci obi oka oca oba okai okei ", 10))
}
