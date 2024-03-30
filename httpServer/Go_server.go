package main

import (
	"fmt"
	"log"
	"net/http"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	for i := 2; i <= int(100000000); i++ {
		if isPrime(i) {
			count++
		}
	}
	fmt.Fprintf(w, "Total prime numbers: %d", count)
	fmt.Printf("Total prime numbers: %d", count)
	fmt.Println()
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/slow", slowHandler)
	http.HandleFunc("/fast", helloWorldHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
