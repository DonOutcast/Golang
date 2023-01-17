package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := readString()
	words := strings.Fields(phrase)
	var abbr []rune

	for _, word := range words {
		letter := []rune(word)[0]
		if !unicode.IsLetter(letter) {
			continue
		}
		abbr = append(abbr, unicode.ToUpper(letter))
	}

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
