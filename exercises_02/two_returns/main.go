package main

import "fmt"

func two_returns(x int) (int, bool) {
	half := x / 2
	even := x%2 == 0
	return half, even
}

func main() {
	fmt.Println(two_returns(1))
	fmt.Println(two_returns(2))
}
