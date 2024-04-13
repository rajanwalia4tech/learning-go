package main

import "fmt"

func main() {

	fmt.Print("Enter a number: ")
	var input float32
	fmt.Scanf("%f", &input) // 12
	output := input * 2
	fmt.Println(output) // 24
}
