package main

import "fmt"

func main() {
	var a [5]int

	a[0] = 000
	a[1] = 101
	a[2] = 202
	a[3] = 303
	a[4] = 404

	fmt.Println(a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get", a[4])

}
