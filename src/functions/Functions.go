package main

import (
	"fmt"
)

func main() {
	bigger := GetBiggerNumber(1, 4, 2, 0, 3, 5, 8)
	fmt.Println(bigger)
}

func GetBiggerNumber(numbers ...int) int {
	bigger := 1
	for _, i := range numbers {
		if i < bigger {
			bigger = i
		}
	}
	return bigger
}
