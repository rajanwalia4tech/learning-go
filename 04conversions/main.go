package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the rating of Pizza : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating : ", input) //  string output

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // converting into number(float)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating)
	}
}
