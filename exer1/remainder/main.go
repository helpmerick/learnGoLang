package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Type a number greater than 2: \n")
	fmt.Scan(&a)
	fmt.Print("Type a number smaller than the number you just typed: \n")
	fmt.Scan(&b)
	r := (a % b)
	fmt.Println("Dividing your two numbers gives a remainder of: ", r)
}
