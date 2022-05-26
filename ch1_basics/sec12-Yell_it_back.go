package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)

	s := msg + strings.Repeat("!", l)
	fmt.Println(s)

	// $ go run ch1_basics/sec12-Yell_it_back.go Hello
	// result: Hello!!!!!

	s = strings.ToUpper(s)
	fmt.Println(s)
	// $ go run ch1_basics/sec12-Yell_it_back.go Hello
	// result: HELLO!!!!!

	m := strings.Repeat("!", l)
	msg = m + msg + m
	fmt.Println(msg)
	// $ go run ch1_basics/sec12-Yell_it_back.go Hello
	// result: !!!!!Hello!!!!!
}
