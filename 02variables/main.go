package main

import "fmt"

// JWT_TOKEN :="skldgjslkgj" // can't use this type of implicit declaration for global variables
var JWT_TOKEN string = "slkjgslkgjsklgj"

func main1() {
	fmt.Println("JWT ---- ", JWT_TOKEN)

	fmt.Println("------ Variables ------")

	var age int // = 34 // Default value for int is 0
	fmt.Printf("Type of age : %T\n", age)
	fmt.Println("VALUE :", age)

	fmt.Println("----------------------")

	var isLoggedIn bool = true // Default value for bool is false
	fmt.Println("Value : ", isLoggedIn)
	fmt.Printf("Type of age : %T\n", isLoggedIn)
	fmt.Printf("VALUE : %D\n", isLoggedIn)

	fmt.Println("----------------------")

	var name string = "RAJAN WALIA" // default value for string is empty string
	fmt.Println("Value : ", name)
	fmt.Printf("Type of age : %T - len %d\n", name, len(name))
	fmt.Printf("VALUE : %D\n", name)

	var gravity float32 = 1233.343434545 // default value for float is 0

	fmt.Println("----------------------")

	fmt.Println("Value : ", gravity)
	fmt.Printf("Type of age : %T\n", gravity)
	fmt.Printf("VALUE : %D\n", gravity)

	fmt.Println("---- Implicit variables ------")
	/*
		Those variable whose type is defined with the value we assign in them and
		we can't put another type value in them
	*/

	college := "MAIMT"

	fmt.Println("College name : ", college)

}
