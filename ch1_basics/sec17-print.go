package main

import "fmt"

func main() {
	var brand string
	fmt.Printf("%q\n", brand) // result: ""

	var brand2 = "Google"
	fmt.Printf("%q\n", brand2) // result: "Google"

	fmt.Printf("%s\n", brand2) // result: Google

}
