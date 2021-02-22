package main

import "fmt"
// (*) star represent a pointer
func main() {
 a := 5
 b := &a
 fmt.Println(a,b)
 fmt.Printf("%T\n", b)

// using * to read val from address
fmt.Println(*b)
fmt.Println(*&a)

}

