package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])
	// fmt.Println(arr[4]) // will cause an error as the array is of length 3

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
}
