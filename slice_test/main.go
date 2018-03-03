package main

import "fmt"

func main() {
	var list1 []int
	fmt.Println("capacity of list1:", cap(list1))
	list2 := []int{1, 2}
	fmt.Println("capacity of list2:", cap(list2))
	list3 := make([]int, 3)
	fmt.Println("capacity of list3:", cap(list3))
}
