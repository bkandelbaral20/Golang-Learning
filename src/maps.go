package main

import "fmt"

func main() {

//  Define map
//	emails := make(map[string]string)
//
////	Assign kv
//  	emails["binjita"] = "binjita@gmail.com"
//	emails["santosh"] = "santosh@gmail.com"
//	emails["sanju"] = "sanju@gmail.com"

//	declare map and rv
	emails := map[string] string { "binjita" : "binjita@gmail.com", "santosh" : "santosh@gmail.com", "sanju" : "sanju@gmail.com"}

//	print it
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["santosh"])

// Delete from map
	delete(emails, "sanju")
	fmt.Println(emails)

}

