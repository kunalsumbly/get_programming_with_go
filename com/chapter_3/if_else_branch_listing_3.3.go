package main

import (
	"fmt"
)

func main() {
	var room = "cave"

	if room == "cave" {
		fmt.Println("You find yourself in a dimly lit cavern")
	} else if room == "entrance" {
		fmt.Println("There is a cavern entrance here and path goes to the west")
	} else if room == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountain")
	} else {
		fmt.Println("Everything is white")
	}
}
