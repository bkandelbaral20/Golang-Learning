package main

import "fmt"

func main() {

	//  long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//	short method
	for i := 1; i<=10; i++ {
		 fmt.Printf("Number is %d\n", i)
	}

	// fizzbuzz
	for i := 1; i <= 20; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
