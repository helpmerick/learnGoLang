package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	var h int
	fmt.Print("Type a number greater than 3: \n")
	fmt.Scan(&h)
	h, even := half(h)
	fmt.Println(h, even)
}
