package main

import (
	"fmt"
)

func main() {
	var userInput string
	fmt.Scan(&userInput)
	switch userInput {
	case "en":
		fmt.Println("English")
	case "fr":
		fmt.Println("French")
	case "ru":
		fmt.Println("Russian")
	case "rus":
		fmt.Println("Russian")
	default:
		fmt.Println("Unknown")
	}
}
