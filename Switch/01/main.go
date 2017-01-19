package main

import "fmt"

func main() {
	switch "Miguel" {
	case "Trudi":
		fmt.Println("Hello querida")
	case "Martin":
		fmt.Println("Hey old man!")
	case "Miguel":
		fmt.Println("Machete!")
	default:
		fmt.Println("Where did everyone go???")
	}
}
