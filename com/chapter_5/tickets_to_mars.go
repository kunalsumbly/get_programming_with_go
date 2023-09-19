package main

import (
	"fmt"
	"math/rand"
)

const departureDateForMars = 13 - 10 - 2020
const distanceToMarsInKms = 62100000

func main() {

	fmt.Println("Spaceline                Days   Trip type   Price")
	fmt.Println("=================================================")

	// declare variable here
	spaceLineText, trip := "", ""

	// iterate over 10 trips
	for count := 1; count <= 10; count++ {

		spaceLineTextSelect := rand.Intn(3)

		switch spaceLineTextSelect {
		case 0:
			spaceLineText = "Virgin Galactic"
		case 1:
			spaceLineText = "Space Adventures"
		case 2:
			spaceLineText = "SpaceX"

		}

		speedOfSpaceShipInKmsPerSec := rand.Intn(15) + 16 // 16-30 km/sec
		daysForTrip := distanceToMarsInKms / (speedOfSpaceShipInKmsPerSec * 24 * 3600)
		// trip cost varies from 36 million to 50 million depending on the speed of the space ship i.e add  20 to the speed of the spaceship
		price := 20 + speedOfSpaceShipInKmsPerSec // in millions

		// return trip logic
		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
			// number of days should also increase as a part of round trip
			daysForTrip = daysForTrip * 2
		} else {
			trip = "One-way"
		}

		// positive number pads spaces to the left and -ve number pads spaces to the right
		fmt.Printf("%-23v %4v %12v $%4v \n", spaceLineText, daysForTrip, trip, price)

	} // closed for loop

}
