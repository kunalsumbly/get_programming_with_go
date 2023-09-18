package main

import "fmt"

/*
Problem statement 1:-  we are calculating the age on mars in earth years for a person living on earth
Problem statement 2 :- we are calculating the weight of a person on mars which is usually 37% of the weight on earth,

	as mars's gravitational pull is weaker compared to earth
*/
func main() {
	const myAgeOnEarth = 38
	const earthDays = 365
	const marsEarthDays = 687
	const myWeightEarth = 75

	// 38 * 365 = y * 687, which represents the time taken to go around the sun, so we need to find out y
	fmt.Println("Age on mars in terms of earth years would be ", myAgeOnEarth*earthDays/marsEarthDays, " years")

	// calculating my weight in kgs on mars
	fmt.Println("My weight on mars would be ", myWeightEarth*.37)

	// formatted print
	fmt.Printf("My weight on the surface of Mars is %v kgs", myWeightEarth*.37)
	fmt.Printf(" and I would be %v years old", myAgeOnEarth*earthDays/marsEarthDays)

}
