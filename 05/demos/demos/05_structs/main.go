package main

import "fmt"

type car struct { // package level struct
	Brand string
	TopSpeed string
}

func main() {

	type user struct { // here the struct can only be used at the main level
		ID int
		FirstName string
		LastName string
	}

	var u user // u will be initialized with it's "zero" values

	fmt.Println(u)

	u.ID = 1
	u.FirstName = "Eduard"
	u.LastName = "Voiculescu"

	fmt.Println(u)

	u2 := user{ ID:2, FirstName: "Denisa", LastName: "Voiculescu"}

	fmt.Println(u2)

	car1 := car{
		Brand: "Honda",
		TopSpeed: "100",
	}

	fmt.Println(car1)

}
