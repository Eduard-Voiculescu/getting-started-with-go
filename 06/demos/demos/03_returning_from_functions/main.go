package main

import (
	"errors"
	"fmt"
)

// _ is a write only variable because in go you have to use every variable that is declared, if not a compilation error will occur

func main() {
	port := 3000
	// we can so call startWebServer(port) without any return value and go will accept it
	fmt.Println("Running startWebServer without any return value assigned to it.")
	startWebServer(port)

	fmt.Println()

	fmt.Println("Runing startWebServer with a return value assigned to it.")
	returnPort, err := startWebServer(port)
	_, err1 := startWebServer(port)
	fmt.Println(returnPort, err)
	fmt.Println(err1)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return port, errors.New("Something went wrong")
}
