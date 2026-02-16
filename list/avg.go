package main

import "fmt"

func main() {
	nums := []int{10, 20, 32, 48, 98, 71}

	
	if len(nums) == 0 {
		fmt.Println("Average: 0")
		return
	}

	sum := 0
	for _, val := range nums {
		sum += val
	}
	avg := float64(sum)/float64(len(nums))
	fmt.Printf("The average is: %.2f\n", avg)
}