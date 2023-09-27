package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)           //0.3333333333333333
	fmt.Printf("%v\n", third)    //0.3333333333333333
	fmt.Printf("%f\n", third)    //0.333333
	fmt.Printf("%.3f\n", third)  // 0.333
	fmt.Printf("%4.2f\n", third) //0.33
	//Note the following:
	//
	// The value third is approximately 0.3333333333333333.
	//	With %4.2f, it's printed as 0.33, following the precision of 2 decimal places.
	// The width 4 is technically met since the entire string 0.33 is 4 characters long, including the decimal point.
	// In the format specifier, such as %4.2f, the numbers before and after the dot (.) represent width and precision respectively.
	//	Width: This is the minimum number of characters the string should take up.
	//	If the value to be printed is fewer characters than this number, it will be padded with spaces (by default) to meet this width.
	//	In your example, 4 is the width. However, if the value itself takes up more space than the specified width, the width will be ignored, and the value will be printed in full.
	// Precision: For the %f format specifier, precision defines the number of digits to be printed after the decimal point.
	//In your example, 2 is the precision, meaning that the output will be formatted to show only two decimal places.
	fmt.Printf("%10.2f\n", third)  //      0.33 space padding on the left to fill the remaining space
	fmt.Printf("%2.2f\n", third)   //      0.33 space padding on the left to fill the remaining space
	fmt.Printf("%010.2f\n", third) //0000000.33  pad left with zeroes
}
