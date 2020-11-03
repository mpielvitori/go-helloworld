package main

import (
	"fmt"
	"time"
)

// func main() {
// 	err := CountDownMultiple(5, "Rocket1", "Rocket2", "Rocket3")
// 	if err != nil {
// 		fmt.Println("Success")
// 	}
// }

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
