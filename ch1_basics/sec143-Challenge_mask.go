package main

import (
	"fmt"
	"os"
)

func main() {
	agrs := os.Args[1:]
	if len(agrs) == 0 {
		fmt.Println("give me some text.")
	}

	text := agrs[0]
	size := len(text)

	word := "http://"
	sizeWord := len(word)

	fmt.Println(text)

	buf := []byte(text)
	isMask := false
	fmt.Println(string(buf))

	for i := 0; i < size; i++ {
		if string(buf[i:i+sizeWord]) == word {
			isMask = true
			i += sizeWord
		}
		c := buf[i]
		switch c {
		case ' ', '\n', '\t':
			isMask = false
		}

		if isMask {
			c = '*'
		}

		buf[i] = c
		// fmt.Println(string(buf[i : i+sizeWord]))
	}
	fmt.Println(string(buf))
}
