package main

import (
	"fmt"
	"math/rand"
)

/*
Generate the dd-mm-yyyy date randomly based on switch case , only the year bit is hardcoded to 2018
*/

// era scope is available throughout the package
var era = "AD"

func main() {
	year := 2018                               // scope of year variable is limited to the main function
	switch month := rand.Intn(12) + 1; month { // month scope if only visible inside switch
	case 2:
		// handle for february
		day := rand.Intn(28) + 1 // scope of day var is limited to case
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 // this is a new day var
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 // this is a new day var
		fmt.Println(era, year, month, day)
	}
}
