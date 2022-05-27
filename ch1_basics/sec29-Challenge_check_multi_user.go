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
	user1     = "jack"
	pass1     = "1888"
	user2     = "n"
	pass2     = "nb"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}
	u := os.Args[1]
	p := os.Args[2]

	if u != user1 && u != user2 {
		fmt.Printf(errAccess, u)
	} else if u == user1 && p != pass1 {
		fmt.Printf(wrongPass, u)
	} else if u == user2 && p != pass2 {
		fmt.Printf(wrongPass, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
