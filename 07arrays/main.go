package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")
	var fruits [4]string

	fruits[0] = "banana"
	fruits[1] = "apple"

	fmt.Println("Arrays ", fruits)               // [banana apple  ]
	fmt.Println("Arrays Length : ", len(fruits)) // 4

	// Initializing while declaring
	var vegetables = [5]string{"potato", "onion", "chilly"}
	fmt.Println(vegetables)      // [potato onion chilly  ]
	fmt.Println(len(vegetables)) // 5
}
