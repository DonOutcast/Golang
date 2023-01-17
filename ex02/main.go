package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	fmt.Scan(&source, &times)
	var result string
	for i := 1; i <= times; i++ {
		result += source
	}
	fmt.Println(result)
}
