package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// var argQty := flag.NArg()

	// fmt.Println("%d parameters", len(argQty))

	//for i:= 0; i < len(flag.Arg)

	// Create program that will greet user by name and take drink order and age
	// If the user orders an alchoholic drink and under 21 then report them to police

	// use maps and slices to have a beautiful menu

	var name string
	fmt.Println("Hello!  What is your name?")
	if _, err := fmt.Scanln(&name); err != nil {
		log.Fatal("Oh no! Bad name! Poop", err)
	}

	var drink string
	fmt.Println("What would you like to drink, " + name + "?")
	if _, err := fmt.Scanln(&drink); err != nil {
		log.Fatal("Oh no! Bad drink! Poop", err)
	}

	fmt.Println("...what is your age?")

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
