package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "carl"

	fmt.Println(len(name))
	// result: 4

	name = "เทส"
	fmt.Println(len(name))
	// result: 9 (9 bytes)

	fmt.Println(utf8.RuneCountInString(name))
	// result: 3
}
