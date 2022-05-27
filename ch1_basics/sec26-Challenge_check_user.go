package main

import (
	"fmt"
	"os"
)

const (
	usage     = "Usage: [username] [password]"
	errAccess = "Access denied for %q\n"
	wrongPass = "Invalid password for %q\n"
	accessOK  = "Access granted to %q\n"
	user      = "jack"
	pass      = "1888"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}
	u := os.Args[1]
	p := os.Args[2]

	if u != user {
		fmt.Printf(errAccess, u)
	} else if p != pass {
		fmt.Printf(wrongPass, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
