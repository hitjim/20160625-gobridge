package main

import "fmt"

type user struct {
	name   string
	logins int
}

func (u user) greet() {
	fmt.Printf("Hello there! My name is %s!\n", u.name)
}

func (u *user) login() {
	u.logins++
}

func main() {
	j := user{name: "Jacob"}

	j.greet()

	fmt.Println("Logins before: ", j.logins)

	j.login()

	fmt.Println("Logins before: ", j.logins)
}
