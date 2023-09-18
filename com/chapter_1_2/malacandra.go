package main

import "fmt"

func main() {
	const distanceToMalacandraInKms = 56000000
	const timeToTravelInDays = 28
	var daysToHours = 28 * 24
	fmt.Printf("speed to cover the distance in kms of %v in number of days %v , in km/h needs to be %v", distanceToMalacandraInKms, timeToTravelInDays, distanceToMalacandraInKms/daysToHours)

}
