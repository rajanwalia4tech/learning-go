package main

import "fmt"

func main() {
	age := 20
	var agePointer *int
	fmt.Println("Value of Pointer is ", agePointer) // default value is nil
	agePointer = &age

	fmt.Println("Value of Pointer is ", agePointer) // memory address in hexadecimal format
	fmt.Println("dereferencing pointer", *agePointer)

	*agePointer = *agePointer + 2
	fmt.Println("dereferencing pointer", age)

	x := &age
	fmt.Println(x)
}
