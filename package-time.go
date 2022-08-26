package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2009, 10, 10, 10, 10, 10, 10, time.UTC)

	fmt.Println(utc)
	layout := "2000-01-21"
	parse, _ := time.Parse(layout, "2020-10-01")

	fmt.Println(parse)
}
