package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	feet := os.Args[1]

	n, err := strconv.ParseFloat(feet, 64)
	if err != nil {
		fmt.Printf("error: %q is not a number.\n", feet)
		return
	}

	m := n * 0.3048
	fmt.Printf("%g feet is %g meters.\n", n, m)
}
