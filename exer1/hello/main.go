package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please type your name: \n")
	fmt.Scan(&name)
	fmt.Println("Hello, ", name)
}
