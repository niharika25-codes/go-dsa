package main

import "fmt"

func get_even_odd(nums []int) ([]int, []int) {
	even := make([]int, 0, len(nums))
	odd := make([]int, 0, len(nums))
	for _, val := range nums {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	return even, odd
}

func main() {
	li := []int{1, 2, 3, 9, 4, 0, 8}
	even, odd := get_even_odd(li)
	fmt.Println("Even: ", even)
	fmt.Println("Odd: ", odd)
}