package main

import "fmt"

//creating functions

// This function name is greeting and take name as a string type and also return string
func greeting(name string) string {
	return "Hello " + name
}

// function that will takes 2 parameter num1 and num2 type of int and also return int
func sumOfTwo(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Binjita"))
	fmt.Println(sumOfTwo(2, 3))
}
