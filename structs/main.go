package main

import "fmt"

// User is a user in our system
type User struct {
	Name   string
	Badge  int
	Salary float64
}

func main() {

	var u User

	fmt.Println(u)

	u.Name = "Jabob"
	u.Salary = 750000
	u.Badge = 3232

	fmt.Println(u)

	kell := User{
		Name:   "Kell",
		Badge:  1002,
		Salary: 1,
	}

	fmt.Println(kell)

	// generally don't do this way
	anna := User{"Anna", 2002, 2323223}
	fmt.Println(anna)
}
