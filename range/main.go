package main

import (
	"fmt"
)

func main() {
	drinks := []string{
		"coke",
		"pepsi",
		"mountain dew",
		"beer",
		"wine",
		"rum",
	}

	// i is index
	for i := range drinks {
		fmt.Println("Would you like some", drinks[i])
	}

	// _ is index d is string
	for _, d := range drinks {
		fmt.Println(d)
	}

	menu := map[string]float32{
		"coke":         3.44,
		"pepsi":        4.52,
		"mountain dew": 3.65,
		"beer":         6.43,
		"wine":         2.33,
		"rum":          2.22,
	}

	for name, cost := range menu {
		fmt.Println(name, "is", cost)
	}
}
