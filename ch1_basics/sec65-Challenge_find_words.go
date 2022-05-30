package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	query := os.Args[1:]
	text := "cat run and run and run and run"
	words := strings.Fields(text)
	for _, v := range query {
		finder := strings.ToLower(v)
		for i, w := range words {
			if finder == w {
				fmt.Printf("#%d : %s \n", i+1, w)
				break
			}
		}
	}
}
