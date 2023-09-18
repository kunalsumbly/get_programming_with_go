package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	if exit {
		fmt.Println("I left the cave")
	} else {
		fmt.Println("I didnt leave the cave")
	}
}
