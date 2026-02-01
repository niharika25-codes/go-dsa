package main

import "fmt"

func linearSearch(arr []int, x int) int {
	for i, n := range arr {
		if n == x {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{3, 5, 7, 9, 11}
	x := 7
	fmt.Println(linearSearch(arr, x))
}