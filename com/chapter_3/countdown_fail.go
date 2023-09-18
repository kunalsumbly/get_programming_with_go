package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*
Generate the logic of launch where in there is a chance of launch fail in 1 out of 100
*/
func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		// generate a launch fail chance in 1 out of 100
		randomNumber := rand.Intn(100)
		if randomNumber == 99 {
			break
		}
		count--
	}

	if count == 0 {
		fmt.Println("lift off !!!")
	} else {
		fmt.Println("Launch failed")
	}

}
