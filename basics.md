
## Using scanf

```
package main

import "fmt"

func main() {
    fmt.Println("Enter the temperature in farenhite : ")
    var tempFahrenheit float64
    fmt.Scanf("%F", &tempFahrenheit)

    celciousTemperature := (tempFahrenheit - 32) * 5 / 9

    fmt.Println("Temperature in Celcious is  : ", celciousTemperature)
}

```


## User Input using bufio

```
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name : ")

	// comma ok || err ok - syntax also kind of try catch

	//  The following syntax will always return string
	input, _ := reader.ReadString('\n') // read till '\n' comes

	fmt.Println("Hello there ", input) // rajan
	fmt.Printf("Type of input is %T", input) // string
}
```

## Converting the userinput into the corrosponding datatype
```
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

```