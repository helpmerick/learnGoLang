package main

import "fmt"

func main() {
	var a int
	fmt.Print("Type a number greater than 3: \n")
	fmt.Scan(&a)
	h := (a / 2)
	fmt.Print("Half of your number is: ", h)
	r := (a % 2)
	if r == 0 {
		fmt.Print("\nYour number was even.")
	} else {
		fmt.Print("\nYour number was odd.")
	}
}
