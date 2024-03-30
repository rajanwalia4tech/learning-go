package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type UserRecord struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	firstName   string `json:"firstName"`
	middleName  string `json:"middleName"`
	lastName    int    `json:"lastName"`
	date        string `json:"date"`
	description string `json:"description"`
	address     string `json:"address"`
	caseNumber  string `json:"caseNumber"`
}

func createShoppingList(data [][]string) []UserRecord {
	// convert csv lines to array of structs
	var userList []UserRecord
	for i, line := range data {
		var record UserRecord
		if i > 0 { // omit header line
			for j, field := range line {
				if j == 11 {
					var name = strings.Split(field, " ")
					record.firstName = name[0]
					if len(name) == 2 {
						// record.lastName = name[1]
					}
					if len(name) == 3 {
						record.middleName = name[1]
						// record.lastName = name[2]
					}
				}
			}
		}
		userList = append(userList, record)
	}
	fmt.Print(userList[0])
	return userList
}

func main() {
	f, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(data)
	createShoppingList(data)
}
