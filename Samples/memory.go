package main

import "fmt"

func main() {

	a := 43
	b := 56
	c := "Rick"

	fmt.Println("a -", a)
	fmt.Println("a's address -", &a)
	fmt.Println("b -", b)
	fmt.Println("b's address -", &b)
	fmt.Println("c -", c)
	fmt.Println("Rick's address -", &c)
}
