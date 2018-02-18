package main

import "fmt"

func main() {
	var sum int
	var i int = 1
	limit := 1000
	for i < limit {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
		i++
	}
	fmt.Println("Sum of all multiples of 3 and 5 between", limit, "is", sum)
}
