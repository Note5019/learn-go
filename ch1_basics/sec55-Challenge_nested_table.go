package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max, _ := strconv.Atoi(os.Args[1])
	max = int(max)

	fmt.Printf("%1s | ", "x")
	for j := 1; j <= max; j++ {
		fmt.Printf("%5d", j)
	}
	fmt.Println()
	for i := 1; i <= max; i++ {
		fmt.Printf("%1d | ", i)
		for j := 1; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}

}
