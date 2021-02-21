package main

import "fmt"

func main() {

	// If else statement
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//	else if statement

	// creating variable to check the statements
		color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Printf("Color is blue")
	} else {
		fmt.Println("Color is not red or blue")
	}

	//	Switch statement
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}

}
