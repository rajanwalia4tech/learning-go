package main

func main() {

	/* 1st way

	var a, b, sum int
	fmt.Println("Enter the number : ")
	fmt.Scanf("%d", &a) // %d for integer, &number for address
	fmt.Scanf("%d", &b)
	sum = a + b
	fmt.Println("You entered: ", sum)

	*/

	/* 2nd Way

	var a, b, sum int
	fmt.Println("Enter the number : ")
	fmt.Scanf("%d %d", &a, &b) // CMT: give the input in spaces
	sum = a + b
	fmt.Println("You entered: ", sum)

	*/

	/* 3rd way

	var a, b, sum int
	fmt.Println("Enter the number : ")
	fmt.Scanln(&a, &b) // CMT: give the input in spaces - here the Scanln will know the type by variable name
	sum = a + b
	fmt.Println("You entered: ", sum)

	*/

	/* bufio.Reader implements the reader interface, allowing you to read byte by byte or line by line from various sources (including standard input).
	// ReadString is a convenient method for reading text until a specific delimiter (like newline).
	reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter your name:")
		name, err := reader.ReadString('\n') // Read until newline

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		fmt.Println("output : ", name[:len(name)-1])

	*/

	/* CMT: Reading String lines using Scanner

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter lines of text (type 'exit' to quit):")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}
		fmt.Println("You entered:", text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	*/

}
