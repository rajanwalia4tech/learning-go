package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guess int
	fmt.Println("Guess a number between 1-100 in 5 tries!!!")
	target := rand.Intn(100) + 1
	success := false
	for attempt := 1; attempt <= 5; attempt++ {
		fmt.Print("Guess a number : ")
		fmt.Scan(&guess)
		if guess < target {
			fmt.Println("Your guess is too Low...")
		} else if guess > target {
			fmt.Println("Your guess is too High...")
		} else {
			fmt.Println("You guessed it right!")
			success = true
			break
		}
	}

	if success == false {
		fmt.Println("You are looser, hahaha, Number was ", target)
	}
}
