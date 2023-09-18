package main

import "fmt"

/*
How long will it take to travel to Mars from earth if we travel at the speed of light
1) speed of light = 299792 km/s
2) distance between mars and earth when they are close to each other = 56,000,000 km
3) distance between mars and earth when they are the farthest i.e on the opposite sides of the sun = 401,000,000 km
*/
func main() {

	const speedOfLightInKmsPerSec = 299792
	// the distance when earth and mars are closest
	var distanceBetweenEarthMars = 56000000
	fmt.Printf("Time taken to cover the distance of %v in kms between earth and mars when travelling at the speed of light , would be %v seconds \n", distanceBetweenEarthMars, distanceBetweenEarthMars/speedOfLightInKmsPerSec)
	// the distance when earth and mars are the farthest
	distanceBetweenEarthMars = 401000000
	fmt.Printf("Time taken to cover the distance of %v in kms between earth and mars when travelling at the speed of light , would be %v seconds \n", distanceBetweenEarthMars, distanceBetweenEarthMars/speedOfLightInKmsPerSec)
}
