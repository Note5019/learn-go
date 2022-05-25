package main

import "fmt"

func main() {
	speed := 100 // speed is int
	force := 2.5 // force is float64

	// speed = speed * int(force)
	// fmt.Println(speed)             // result: 200 (wrong)
	// fmt.Println(force, int(force)) // result: 2.5, 2

	speed = int(float64(speed) * force)
	fmt.Println(speed) // result: 250 (correct !)

}
