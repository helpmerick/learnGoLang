package main

import "fmt"

func main() {
	fmt.Println(greet("Trudi ", "Hoffman"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
