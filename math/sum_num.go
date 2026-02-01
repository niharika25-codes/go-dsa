package main

import "fmt"

func find_sum(n int64) int64 {
	if n%2 == 0 {
		return (n / 2) * (n + 1)
	}
	return ((n + 1) / 2) * n
	
}

func main() {
	var n int64
	fmt.Println("Enter a number")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
	fmt.Printf("sum of first %d numbers is %d", n, find_sum(n))
}