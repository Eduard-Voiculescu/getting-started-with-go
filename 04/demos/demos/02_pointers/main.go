package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)

	*firstName = "Arthur"
	fmt.Println(firstName) // will print out the memory address (value that the pointer is holding)

	fmt.Println(*firstName) // will print out the value at the memory address
}
