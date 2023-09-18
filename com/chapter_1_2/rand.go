package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// generates a random number between 1-10
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	// generate a random number between 56000000 to 401000000 (distance between earth and mars)
	// first find the difference between ,401000000 - 56000000 = 345000000
	// and then use it in the -> rand.Intn(345000001) to generate random values between 0 and 345000000 and then we need to add 56000000 to the number generated
	var distance = rand.Intn(345000001) + 56000000
	fmt.Println(distance)

}
