package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port int, numberOfRetries int) {
	println("Starting server...")
	// do important things
	println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}

/*
// go will accept not putting int after port as it will take to be an int, the same as numberOfRetries
func startWebServer(port, numberOfRetries int) {
	println("Starting server...")
	// do important things
	println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}
*/
