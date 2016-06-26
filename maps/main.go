package main

import "fmt"

func main() {
	scores := make(map[string]int)

	scores["jacob"] = 75
	scores["anna"] = 82
	scores["kell"] = 66
	scores["carter"] = 44
	scores["jimrice"] = 1000000000

	fmt.Println(scores)

	j := scores["jacob"]
	fmt.Println("Jacob's score", j)

	k := scores["katie"]
	fmt.Println("katie's score:", k)

	if _, ok := scores["liam"]; !ok {
		fmt.Println("Liam had no score!")
	}

	delete(scores, "jimrice")

	fmt.Println(scores)

	scores2 := map[string]int{"a": 100, "b": 200, "c": 300}
	fmt.Println(scores2)
}
