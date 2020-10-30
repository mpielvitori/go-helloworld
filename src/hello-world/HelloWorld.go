package main

import (
	"fmt"
)

var (
	name = "Martin"
)

func Hello() string {
	greeting := "Hello"
	return greeting + " " + name + "!"
}

func main() {
	fmt.Println(Hello())
}
