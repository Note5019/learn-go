package main

import (
	"fmt"
	"os"
)

// var Args []string
// ex: go run file.go hi yo
// Args[0] => path to the program
// Args[1] => "hi"
// Args[2] => "yo"

func main() {
	// way 1
	// fmt.Printf("%#v\n", os.Args)

	// $ go run ch1_basics/sec6-Get_input_from_terminal.go
	// result: [/var/folders/3p/7d_1qt7x2qv1vv2x_clvkcs00000gn/T/go-build3571670234/b001/exe/sec6-Get_input_from_terminal]

	// $ go build -o greeter  // << Go will build&compile program to file named greeter
	// $ ./greeter
	// result: []string{"./greeter"}
	// $ ./greeter hi hellow hep
	//result: []string{"./greeter", "hi", "hellow", "hep"}

	// way 2
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st argument: ", os.Args[1])
	fmt.Println("2nd argument: ", os.Args[2])
	fmt.Println("3th argument: ", os.Args[3])
	fmt.Println("len Args: ", len(os.Args))
	// $ go build -o greeter
	// $ ./greeter hi hellow hep
	// result:
	// []string{"./greeter", "hi", "hellow", "hep"}
	// Path:  ./greeter
	// 1st argument:  hi
	// 2nd argument:  hellow
	// 3th argument:  hep
	// len Args:  4

}
