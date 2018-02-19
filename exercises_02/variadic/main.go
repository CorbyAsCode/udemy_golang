package main

import "fmt"

func maximum(x ...int) int {
	var a int
	for _, i := range x {
		if i > a {
			a = i
		}
	}
	return a
}

func main() {
	fmt.Println(maximum(55, 4, 9, 23))
}
