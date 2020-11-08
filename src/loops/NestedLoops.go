package main

import (
	"fmt"
	"time"
)

func CountDownMultiple(from int, rockets ...string) error {
startBreak:
	for _, rocket := range rockets {
		fmt.Println("Launching " + rocket)
		for i := from; i >= 0; i-- {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
			if rocket == "Rocket2" {
				fmt.Println("Aborting!!!")
				break startBreak
			}
		}
		fmt.Println("Boom!")
	}

	return nil
}
