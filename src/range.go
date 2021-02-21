package main

import "fmt"

func main() {

	ids := []int{1, 2, 9, 4, 5, 3}

	// loops through ids using range
	for i, id := range ids {
		fmt.Printf("%d - ID:  %d\n", i, id)
	}
	// Not using index
	for _, id := range ids { //if you not gonna use that i just put underscore
		fmt.Printf("ID : %d\n", id)
	}
	//	Adding Ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum is", sum)

	// range with map
	emails := map[string]string{"binjita": "binjita@gmail.com", "santosh": "santosh@gmail.com", "sanju": "sanju@gmail.com"}
	for k, v := range emails { // kv = key values
		fmt.Printf("%s: %s\n", k, v)
	}
}
