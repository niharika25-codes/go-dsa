package main

import "fmt"

func removeDuplicates(nums []int) int {
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[res-1] != nums[i] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

func main() {
	nums := []int{10, 20, 20, 30, 30, 30, 30}
	fmt.Println(removeDuplicates(nums))
}

func dedupInPlace[S ~[]E, E comparable](s S) S {
	seen := make(map[E]struct{})
	res := 0 // Our write-pointer

	for i := 0; i < len(s); i++ {
		if _, ok := seen[s[i]]; !ok {
			seen[s[i]] = struct{}{}
			s[res] = s[i] // Move unique element to the front
			res++
		}
	}
	return s[:res]
}