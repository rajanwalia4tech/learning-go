package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name : ")

	// comma ok || err ok - syntax also kind of try catch

	//  The following syntax will always return string
	input, error := reader.ReadString('\n') // read till '\n' comes
	if error != nil {
		log.Fatal(error)
	}
	input = strings.TrimSpace(input)
	grade, error := strconv.ParseFloat(input, 64)
	if grade >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// fmt.Println("Hello there ", input)
	// fmt.Printf("Type of input is %T", input)
}
