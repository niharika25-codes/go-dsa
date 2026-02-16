package main

import (
	"fmt"
	"slices"
)

func isSorted(b []int) bool {
	for i := 1; i < len(b); i++ {
		if b[i] < b[i-1] {
			return false
		}
	}
	return true
}

func main() {
	l1 := []int{10, 20, 30, 40}
	l2 := []int{20, 20, 78, 98, 99, 97}
	fmt.Println("Using slices.IsSorted:")
	fmt.Println(slices.IsSorted(l1))
	fmt.Println(slices.IsSorted(l2))
	fmt.Println("using custom isSort:")
	fmt.Println(isSorted(l1))
	fmt.Println(isSorted(l2))
}