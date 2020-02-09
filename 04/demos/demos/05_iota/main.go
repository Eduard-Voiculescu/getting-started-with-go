package main

import "fmt"

const (
	first = iota + 6
	second
)

const (
	third = iota
	fourth
)

const (
	fifth = iota + 6
	sixth = 2 << iota // iota = 1 (01 in bit) and we bit shift to the left for iota so we get (01 -> 10 = 2)
)

// iota can be very useful with many constants
// it is a constant expression
// iota resets between constant blocks
// very useful if you want to create family of constants and make usefulness of iota

// constant expressions are evaluated at compilation time and not runtime so they have to be valid

// iota resets between constant blocks, it is very useful if you need constant blocks of variables to be unique and distintive
func main() {
	fmt.Println(first, second, third, fourth)
	fmt.Println(fifth, sixth)
}
