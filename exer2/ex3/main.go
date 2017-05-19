package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {  //underscore = an index of list (slice)
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(4, 440, 1500, 7045, 1966, 3)
	fmt.Println(greatest)
}
