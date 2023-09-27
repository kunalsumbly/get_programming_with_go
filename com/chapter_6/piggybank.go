package main

import "fmt"

func main() {
	// golang assumes the piggyBank type as float64
	// default type assigned to a float value is float64
	piggyBank := 0.0

	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}

	fmt.Printf("final piggy bank balance, = %v\n", piggyBank)
	fmt.Printf("final piggy bank balance, = %5.2f\n", piggyBank)
}
