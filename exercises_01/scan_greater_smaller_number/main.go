package main

import "fmt"

func main() {
	var small, large int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Enter a large number: ")
	fmt.Scan(&large)
	fmt.Println("Remainder of", large, "divided by", small, "is", large%small)
}
