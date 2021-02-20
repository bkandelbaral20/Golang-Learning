package main
import "fmt"

/*
Main types are :
string
bool
int
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr (non-negative number)
byte - alias of uint8
rune - alias of uint32
float32 float 64
complex64 complex28( really large numbers)

 */

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