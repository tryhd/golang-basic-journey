package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(3.3))
	fmt.Println(math.Round(3.6))
	fmt.Println(math.Floor(3.6))
	fmt.Println(math.Ceil(3.1))

	fmt.Println(math.Max(3.1, 2))
	fmt.Println(math.Min(3.1, 10))

}
