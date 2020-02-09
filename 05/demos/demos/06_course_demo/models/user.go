package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// variable block
var (
	users  []*User // collection of users, will be a slice that contains pointers to the users
	nextID = 1
)
