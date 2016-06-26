package main

import (
	"fmt"
	"log"
)

func main() {
	// Create program that will greet user by name and take drink order and age
	// If the user orders an alchoholic drink and under 21 then report them to police

	// use maps and slices to have a beautiful menu

	// use structs to define the different drinks in menu

	// on drink type - add method print to drink.  Use this method to print menu

	type drink struct {
		name     string
		price    float32
		alchohol bool
	}

	// Can omit drink on type literal b/c of inherited slice literal
	// menu := []drink{
	// 	drink{"beer", 4.55, true},
	// 	drink{"wine", 2.54, true},
	// 	drink{"liquor", 2.56, true},
	// 	drink{"moonshine", 8.99, true},
	// 	drink{"rum", 6.77, true},
	// }
	menu := []drink{
		{"beer", 4.55, true},
		{"wine", 2.54, true},
		{"liquor", 2.56, true},
		{"moonshine", 8.99, true},
		{"rum", 6.77, true},
	}

	fmt.Println("Here is our menu")
	for i, item := range menu {
		fmt.Printf("%d: %s - $%.2f\n", i, item.name, item.price)
	}

	var name string
	fmt.Println("Hello!  What is your name?")
	if _, err := fmt.Scanln(&name); err != nil {
		log.Fatal("Oh no! Bad name! Poop\n", err)
	}

	var choice int32
	fmt.Println("Enter a number to choose your drink!")
	if _, err := fmt.Scanln(&choice); err != nil {
		log.Fatal("Oh no! Bad choice! Poop\n", err)
	}

	// var drink string
	// fmt.Println("What would you like to drink, " + name + "?")
	// if _, err := fmt.Scanln(&drink); err != nil {
	// 	log.Fatal("Oh no! Bad drink! Poop\n", err)
	// }

	var age int32
	fmt.Println("... what is your age?")
	if _, err := fmt.Scanln(&age); err != nil {
		log.Fatal("Oh no! Bad age! Poop\n", err)
	}

	fmt.Println("Next year you will be", age+1)

	if menu[choice].alchohol && age < 21 {
		fmt.Println("Go to jail by way of police.  I'll call them eventually, but if you wouldn't mind, please just turn yourself in.  I'm kind of tired.")
	} else {
		fmt.Println("Hey " + name + " enjoy your " + menu[choice].name)
	}

	// switch drink {
	// case "beer", "wine", "liquor", "whiskey", "moonshine", "rum":
	// 	if age < 21 {
	// 		fmt.Println("Go to jail by way of police.  I'll call them eventually, but if you wouldn't mind, please just turn yourself in.  I'm kind of tired.")
	// 		break
	// 	}

	// 	fmt.Println("Enjoy your", drink, "responsibly"+name)
	// default:
	// 	fmt.Println("Hey "+name+", you are", age, "so enjoy your "+drink+".")
	// }
}
