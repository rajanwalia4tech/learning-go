package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time of Golang")

	currentTime := time.Now()
	fmt.Println(currentTime)

	// The following are the hard coded time formats
	fmt.Println("DateOnly : ", currentTime.Format("2006-01-02")) // Formatting date

	fmt.Println("DateTime Only : ", currentTime.Format("2006-01-02 15:04:05")) // Formatting date

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("createdDte : ", createdDate)

}
