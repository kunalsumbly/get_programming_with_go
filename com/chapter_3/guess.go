package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myNumber = 45
	guessMyNumber(myNumber)
}

func guessMyNumber(myNumber int) {
	// infinite for loop usage
	for {
		fmt.Println("start guessing")
		guessedNumber := rand.Intn(100) + 1
		fmt.Printf("number guessed %v is bigger than my number %v \n", guessedNumber, guessedNumber > myNumber)
		if guessedNumber == myNumber {
			fmt.Printf("Gotcha, I guessed the number %v correctly \n", myNumber)
			break
		} else {
			continue
		}
	}
}
