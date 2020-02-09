package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	// slice will be pointing to the array arr so changing the
	// values of either arr or slice
	// will "change" for both as it is pointing at the same place in 
	// memory, it is kind of a pointer pointing to the arr 
	slice := arr[:] // from the beginning to the end of the array (inclusive:exclusive)

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice) 
	array := [4]int{0,0,0,0}
	coucou := array[0:1]

	fmt.Println(array, coucou)

}
