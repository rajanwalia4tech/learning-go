package main

func main() {
	/*
		for i := 0; i < 5; i++ {
			fmt.Println("Number:", i)
		}
	*/

	/* CMT: Infinite loop

	for {
	    fmt.Println("This loop will run forever!")
	    // Use break to exit the loop if needed
	}

	*/

	/* CMT: Iterate over array
	This is a concise way to iterate over elements in a collection like slices, arrays,
	maps, strings, or channels. The range keyword extracts the element (key-value pair for maps)
	and assigns it to a variable within the loop body.

		names := []string{"Alice", "Bob", "Charlie"}

		for idx, value := range names {
			fmt.Println("idx", idx, "value", value)
		}
	*/
	/* CMT: Iterate over maps
	// Create a map of string keys and int values
	fruits := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}

	// Loop through the map using the range keyword
	for key, value := range fruits {
		fmt.Println("Fruit:", key, ", Count:", value)
	}
	*/

	/* iterate over strings
	message := "Hello, world!"

	for i := 0; i < len(message); i++ {
		fmt.Println(i, string(message[i])) // Print index and character
	}
	*/

	/* iterate over string
	message := "Hello, world!"

	for _, char := range message {
		fmt.Println(char, string(char)) // Print rune value and character
	}
	*/

	/* CMT: Iterate over 2d array
			// Define a 2D array (slice of slices)
	  matrix := [][]int{
	    {1, 2, 3},
	    {4, 5, 6},
	    {7, 8, 9},
	  }

	  // Nested loops for rows and columns
	  for i := 0; i < len(matrix); i++ {
	    for j := 0; j < len(matrix[i]); j++ {
	      fmt.Println("Element at", i, ",", j, ":", matrix[i][j])
	    }
	  }

	*/

	/* CMT: Iterate over 2D array using range
		// Define a 2D array (slice of slices)
	  matrix := [][]int{
	    {1, 2, 3},
	    {4, 5, 6},
	    {7, 8, 9},
	  }

	  // Single loop with range iterating over inner slices
	  for _, row := range matrix {
	    for _, value := range row {
	      fmt.Println(value)  // Print element only
	    }
	  }
	*/

}
