package main

import "fmt"

const FtoC float64 = 32

func main() {
	var farn float64
	fmt.Print("Enter temperature in Farenheit:")
	fmt.Scan(&farn)
	celc := (farn - FtoC) * 5 / 9
	fmt.Println(farn, "Farenheit temp is: ", celc, "Celcius temp is")
}
