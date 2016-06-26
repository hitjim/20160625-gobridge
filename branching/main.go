package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("Even")

	} else {
		fmt.Println("Odd")
	}

	// age is only scoped to this if/else block
	if age := 21 + 33; age > 44 {
		fmt.Println("Your age is greater than 44")
	} else {
		fmt.Println("Your age is uder 44")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
