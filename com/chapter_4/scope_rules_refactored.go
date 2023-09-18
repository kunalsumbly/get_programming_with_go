package main

import (
	"fmt"
	"math/rand"
)

/*
Generate the dd-mm-yyyy date randomly based on switch case , only the year bit is hardcoded to 2018
*/

// era scope is available throughout the package
var era1 = "AD"

func main() {
	for count := 1; count <= 10; count++ {
		year := rand.Intn(1000) + 2000
		generateDate(year)
	}
}

func generateDate(year int) {
	month := rand.Intn(12) + 1
	daysInMonths := 31 // default case
	switch month {     // month scope if only visible inside switch

	case 2:
		// handle for february
		if year%400 == 0 {
			fmt.Println("******************************** we are in a leap year ******************************************")
			daysInMonths = 29

		} else {
			daysInMonths = 28

		}

	case 4, 6, 9, 11:
		daysInMonths = 30
	}

	days := rand.Intn(daysInMonths) + 1
	fmt.Println(era1, year, month, days)
}
