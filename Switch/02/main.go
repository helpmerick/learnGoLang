package main

import "fmt"

func main() {
	switch "Martin" {
	case "Trudi":
		fmt.Println("Hello querida")
		fallthrough
	case "Martin":
		fmt.Println("Hey old man!")
		fallthrough
	case "Miguel":
		fmt.Println("Machete!")
	default:
		fmt.Println("Where did everyone go???")
	}
}
