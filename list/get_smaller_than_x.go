package main

import "fmt"

func get_smaller(nums []int, x int) []int {
	res := make([]int, 0, len(nums))
	for _, val := range nums {
		if val < x {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	nums := []int{8, 100, 20, 40, 3, 7}
	 x := 10
	fmt.Println(get_smaller(nums, x))
}