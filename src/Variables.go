package main
import "fmt"

func main() {
	//  using var

	// var name = "Binjita"
	var age = 23
	const isName = true

//	using shorthand methods
 name := "Binjita"
 size := 1.33
 email := "binjita@gmail.com"

// we can combine two shorthand variables like this
	name, email = "Binjita", "binjita@gmail.com"

// to printout
	fmt.Println(name)
	fmt.Println(email)

	fmt.Println(name,email)

//	to checkout the data type
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size) // float64, we have to implicitly put the float32 or 64 which we want to use
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isName)
}