package main

import (
	"fmt"
	"path"
)

func main() {
	// example 1
	var dir, file string
	dir, file = path.Split("css/main.css")
	dir, file = path.Split("css/main.css")
	fmt.Println("file:", file)
	fmt.Println("dir:", dir)

	// example 2
	// var file string
	// _, file = path.Split("css/main.css")
	// fmt.Println("file:", file)
}
