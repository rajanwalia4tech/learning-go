package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(start, end int) int {
	count := 0
	for num := start; num <= end; num++ {
		if isPrime(num) {
			count++
		}
	}
	return count
}

func main() {
	start := time.Now()
	count := countPrimes(1, int(math.Pow(10, 8)))
	end := time.Now()
	fmt.Println()
	fmt.Println("---------- GO Performance -----------")
	fmt.Println()
	fmt.Printf("Number of prime numbers between 1 and 1 million: %d\n", count)
	fmt.Printf("Time taken: %s\n", end.Sub(start))
	fmt.Println()
}
