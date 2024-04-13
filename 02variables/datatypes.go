package main

import "fmt"

// test := "test" // This syntax doesn't work in global variable
func main() {
	fmt.Println("Length of String :  ", len("rajan")) // 5
	fmt.Println("Hello " + "   World")                // concatenation
	var x string = "Raj"
	fmt.Println(x)
	x = "Nitin"
	fmt.Println(x)
	var address = "delhi"
	fmt.Println(address)
	age := 32
	println(age)

	const pi float32 = 3.1454223 // const values can't be change later
	println(pi)
	const pi2 float32 = 3.1454223 // const values can't be change later
	println(pi == pi2)            //  true

	// Defining multiple variables

	var (
		a = 5
		b = 10
		c = 15
	)
	println(a + b + c) // 30
}
