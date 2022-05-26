package main

import "fmt"

func main() {
	const (
		mon = iota
		tue
		wed
		thu
		fri
		sat
		sun
	)
	fmt.Println(mon, tue, wed, thu, fri, sat, sun) // result: 0 1 2 3 4 5 6

	const (
		EST = -(5 + iota)
		_   // blank-identifier
		MST
		PST
	)

	fmt.Println(EST, MST, PST) // result: -5 -7 -8
}
