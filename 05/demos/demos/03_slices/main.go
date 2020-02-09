package main

import "fmt"

func main() {

	slice := []int{1, 2, 3} // no size specified, we let the compiler manage it so it can be added
	fmt.Println(slice)

	// a slice is not a fixed size entity as an array
	// we can add to it

	// slice2 := [2]int{1, 2} // won't work as we set a specific size to the slice (can't append)
	slice2 := []int{1, 2}
	fmt.Println(slice2)
	slice2 = append(slice2, 4, 42, 27) 
	fmt.Println(slice2)

	slice = append(slice, 4, 42, 27)
	fmt.Println(slice)

	s2 := slice[1:] // take it from index 1 to the end
	s3 := slice[:2] // start at the beginning (0) and get up to (while not including) index 2
	s4 := slice[1:2]

	fmt.Println(slice, s2, s3, s4)
}
