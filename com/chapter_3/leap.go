package main

import (
	"fmt"
	"os"
	"strconv"
)

/* what is a leap year??
<ol>
	<li> if a year is evenly divisible by 4 but not divisible by 100 </li>
	<li> or , if a year is divisible by 400 </li>
</ol>
*/

func main() {
	inputYear, err := extractArgs()

	if err != nil {
		fmt.Println("we have an error while parsing string to int", err.Error())
		os.Exit(1)
	}

	if (inputYear % 400) == 0 {
		fmt.Println("look before you leap as we are in the leap year")
	} else {
		fmt.Println("Keep your feet on the ground as we are not in a leap year")
	}

}

/*
This method will extract input args passed to cli
*/
func extractArgs() (int, error) {
	var inputYear = os.Args[1:]
	// change string to int using strconv package
	year, err := strconv.Atoi(inputYear[0])
	return year, err
}
