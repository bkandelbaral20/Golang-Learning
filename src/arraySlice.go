package main

//fmt is same as "sout" in java, just to print stuff out
import "fmt"

func main() {

	// 1.

	// Declaring Arrays
	var fruitArray [3]string

	//Assigning values
	fruitArray[0] = "Mango"
	fruitArray[1] = "Orange"
	fruitArray[2] = "Banana"

	//2.

	//Using shorthand to declare and assign values in an array
	fruitArray1 := [3]string{"Mango", "Banana", "Oranges"}        //must have 3 element in an array because it is already declared
	fruitSlice := []string{"Mango", "Banana", "Oranges", "Apple"} //can have unlimited elements, because there is no any limit in an array

	// To print
	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])
	fmt.Println(fruitArray1)
	fmt.Println(fruitSlice)

	// To get the total length of an array
	fmt.Println(len(fruitSlice))

	// To get range
	fmt.Println(fruitSlice[2:3]) //start at 1 and stop before 2

}
