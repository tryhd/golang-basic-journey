package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	//with switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case float64:
		fmt.Println("Float", value)
	default:
		fmt.Println("Unknown", value)
	}
}
