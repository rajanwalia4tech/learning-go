
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