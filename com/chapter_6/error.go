package main

import (
	"fmt"
	"math"
)

func main() {
	third := 1.0 / 3.0
	// what does it print ??  mathematically 1/3 +1/3 +1/3 = 0.99
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	// prints ugly 0.30000000000000004
	fmt.Println(piggyBank)
	// dealing with floating-points is not the best choice for representing money, because they cause rounding off errors

	// not a good idea to compare float values directly , instead determine the absolute difference between the two numbers and then ensure that the difference is not too big
	fmt.Println("Compare float values. Are the values equal ? ", piggyBank == 0.3)

	// since the absolute comparisons of the float points is not possible, if we have to compare them anyways then it is better to define a tolerance level (difference between what we expect and the actual value) that we are comfortable with to compare two float values
	toleranceLevel := 0.0001
	fmt.Println("Compare two float values indirectly using a tolerance level, are the values equal ? ", math.Abs(piggyBank-0.3) < toleranceLevel)

	// order of operations multiplication over division and vice-versa , impacts the final results
	calculateCelsiusToFarenheit()
}

func calculateCelsiusToFarenheit() {
	celsius := 21.0
	fmt.Println((celsius/5.0*9.0)+32, "is the temperature in farenheit \n")
	fmt.Println((9.0/5.0*celsius)+32, "is the temperature in farenheit \n")
	// both the above statements print 69.80000000000001 because we are performing division before multiplication
	// to minimise rounding errors , it is recommended that we perform multiplication before division, as shown below
	fmt.Println((celsius*9.0/5.0)+32, "is the temperature in farenheit \n")
	// the above statement prints 69.8
}
