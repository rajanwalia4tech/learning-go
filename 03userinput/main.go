package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name : ")

	// comma ok || err ok - syntax also kind of try catch

	//  The following syntax will always return string
	input, _ := reader.ReadString('\n') // read till '\n' comes

	fmt.Println("Hello there ", input)
	fmt.Printf("Type of input is %T", input)
}
