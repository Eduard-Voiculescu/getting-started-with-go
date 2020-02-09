package main

import "fmt"

func main() {
	const pi = 3.1415
	fmt.Println(pi)
	fmt.Println(pi)
	// pi = 1.2 // error as we can't assign a new value to a constant

	const c = 3
	fmt.Println(c + 1.2) // will work because the interpreter will assign the type of c to a float

	const b int = 4

	//fmt.Println(b + 1.1) // will not work as b is an integer

	// a bunch of code

	fmt.Println(c + 5)
}
