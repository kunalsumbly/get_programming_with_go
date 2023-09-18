package main

import (
	"fmt"
	"os"
)

func main() {
	/* TODO
	https://gobyexample.com/command-line-arguments
	Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
		argsWithProg := os.Args
	    argsWithoutProg := os.Args[1:]
	You can get individual args with normal indexing
	*/

	var room = os.Args[1:]

	if room[0] == "cave" {
		fmt.Println("You find yourself in a dimly lit cavern")
	} else if room[0] == "entrance" {
		fmt.Println("There is a cavern entrance here and path goes to the west")
	} else if room[0] == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountain")
	} else {
		fmt.Println("Everything is white")
	}
}
