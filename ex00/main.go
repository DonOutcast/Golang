package main

import (
	"fmt"
	"time"
)

func main() {
	var user_input string
	fmt.Scanf("%s", &user_input)
	m, _ := time.ParseDuration(user_input)
	fmt.Printf("%s = %.0f min", user_input, m.Minutes())
}
