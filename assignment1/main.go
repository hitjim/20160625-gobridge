package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// var argQty := flag.NArg()

	// fmt.Println("%d parameters", len(argQty))

	//for i:= 0; i < len(flag.Arg)

	// Create program that will greet user by name and take drink order and age
	// If the user orders an alchoholic drink and under 21 then report them to police

	//fmt.Println("Hello!  What is your name?")
	flag.Parse()
	name := flag.Arg(0)

	//fmt.Println("...would you like water or beer?")
	drink := flag.Arg(1)

	//fmt.Println("...what is your age?")
	age, err := strconv.Atoi(flag.Arg(2))

	if err != nil {
		fmt.Println("Age was not a number!")
	} else {
		fmt.Println("Next year you will be", age+1)
	}

	switch drink {
	case "beer", "wine", "liquor", "whiskey", "moonshine", "rum":
		if age < 21 {
			fmt.Println("Go to jail by way of police.  I'll call them eventually, but if you wouldn't mind, please just turn yourself in.  I'm kind of tired.")
		}
	default:
		fmt.Println("Hey "+name+", you are", age, "so enjoy your "+drink)

	}
}
