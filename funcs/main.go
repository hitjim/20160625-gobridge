package main

import "fmt"

func main() {
	s := sum(1, 3)
	fmt.Println(s)

	fmt.Println("4 plus 2 is", sum(4, 2))

	q, r := div(6, 2)
	fmt.Println("6 div 2 is", q, "with remainder", r)
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
