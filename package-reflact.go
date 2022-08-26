package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Nama string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name string
	Desc string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Oki"}
	sampleType := reflect.TypeOf(sample)
	stuctField := sampleType.Field(0)
	required := stuctField.Tag.Get("required")

	fmt.Println(required)

	sample.Nama = ""
	fmt.Println(IsValid(sample))
}
