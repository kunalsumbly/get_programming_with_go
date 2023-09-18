package main

import "fmt"

func main() {
	const speedOfTravelInKmPerHour = 100800
	const distanceBetweenEarthMarsInKms = 96300000

	fmt.Printf("Number of days to travel the distance of %v in kms , at this speed %v , would be %v \n", distanceBetweenEarthMarsInKms, speedOfTravelInKmPerHour, distanceBetweenEarthMarsInKms/speedOfTravelInKmPerHour)

	fmt.Println("days taken=", 955/24)
}
