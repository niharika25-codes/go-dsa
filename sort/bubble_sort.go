package main

import "fmt"

func less(val1, val2 int) bool {
	return val1 < val2
}

func greater(val1, val2 int) bool {
	return val1 > val2
}

func bubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func bubbleSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := true
	for i := 0; i < size-1 && swapped; i++ {
		swapped = false
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = true
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Before sort: ", arr)
	bubbleSort(arr, greater)
	bubbleSort2(arr, less)
	fmt.Println("After sort: ", arr)
}