package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
)

/*
	nickels $0.05
	dimes $0.10
	quarters $0.25

write a program that randomly places nickels, dimes and quarters into an empty piggy bank until the balance contains at least $20.
display the running balance of the piggy bank after each deposit, formatting it with appropriate width and precision
*/
func main() {
	piggyBank := 0.0
	//piggyBank = addToPiggyBankUsingTolerance(piggyBank)

	for piggyBank < 20.00 {
		moneyLabel := rand.Intn(3)

		if moneyLabel == 0 {
			piggyBank += 0.05
			fmt.Printf("added nickel to the piggyBank , sum in piggyBank = %2.2f\n", piggyBank)
		} else if moneyLabel == 1 {
			piggyBank += 0.10
			fmt.Printf("added dime to the piggyBank , sum in piggyBank = %2.2f\n", piggyBank)
		} else {
			piggyBank += 0.25
			fmt.Printf("added quarter to the piggyBank , sum in piggyBank = %2.2f\n", piggyBank)
		}

	}

	fmt.Printf("Final balance in the piggyBank = %2.2f", piggyBank)

}

func addToPiggyBankUsingTolerance(piggyBank float64) float64 {
	tolerance := 0.001

	// why is this needed here ? we are not doing equality check
	for (20.00 - piggyBank) >= tolerance {
		moneyLabel := rand.Intn(3)

		if moneyLabel == 0 {
			piggyBank += 0.05
			fmt.Printf("added nickel to the piggyBank , sum in piggyBank = %2.2f\n", piggyBank)
		} else if moneyLabel == 1 {
			piggyBank += 0.10
			fmt.Printf("added dime to the piggyBank , sum in piggyBank = %2.2f\n", piggyBank)
		} else {
			piggyBank += 0.25
			fmt.Printf("added quarter to the piggyBank , sum in piggyBank = %2.2f\n", piggyBank)
		}

	}
	return piggyBank
}
