package main

import "fmt"

func main() {
	i := 101
	fmt.Println(i)

	increment(&i)

	fmt.Println(i)
}

func increment(n *int) {
	*n++
}
