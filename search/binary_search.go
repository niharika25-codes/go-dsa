package main

import "fmt"

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] > x {
			high = mid - 1
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 10
	fmt.Println(binarySearch(arr, x))
}