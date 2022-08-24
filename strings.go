package main
import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Oki Iskandar","Oki"))
	fmt.Println(strings.Contains("Oki Iskandar","Budi"))

	fmt.Println(strings.Split("Oki Iskandar"," "))
	fmt.Println(strings.ToLower("Oki Iskandar"))
	fmt.Println(strings.ToUpper("Oki Iskandar"))
	fmt.Println(strings.ToTitle("Oki Iskandar"))
	fmt.Println(strings.Trim("   Oki Iskandar    ", " "))
	fmt.Println(strings.ReplaceAll("Oki Oki Iskandar", "Oki", "Budi"))
}