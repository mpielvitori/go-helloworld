package main

import (
	"fmt"
	"time"
)

func CountDown(from int) error {
	for i := from; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	fmt.Println("Boom!")
	return nil
}
