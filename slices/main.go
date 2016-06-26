package main

import "fmt"

func main() {
	s := make([]string, 3)

	fmt.Println(s)

	s[0] = "Apple"
	s[1] = "Banana"
	s[2] = "Cherry"

	fmt.Println("My slice is", len(s), "long")

	fmt.Println(s)

	s = append(s, "Date", "Elderberry", "Fruit")

	fmt.Println("My slice is", len(s), "long")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("The copy is", c)

	sub := s[2:5]

	fmt.Println("subslice is", sub)

	sub[2] = "ICE COLD"
	fmt.Println("slice is", s)
	fmt.Println("The copy is", c)
	fmt.Println("subslice is", sub)

}
