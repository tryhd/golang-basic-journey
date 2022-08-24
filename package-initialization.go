package main

import (
	"fmt"
	"project-pertama/Learn/src/Golang-journey/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
