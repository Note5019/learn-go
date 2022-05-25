package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st argument: ", os.Args[1])
	fmt.Println("2nd argument: ", os.Args[2])
	fmt.Println("3th argument: ", os.Args[3])
	fmt.Println("len Args: ", len(os.Args))
}
