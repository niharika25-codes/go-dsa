package main

import "fmt"

func get_max(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	nums := []int{20, 10, 20, 4, 100}
	fmt.Println("max: ", get_max(nums))
}