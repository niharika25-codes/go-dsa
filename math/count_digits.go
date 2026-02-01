package main

import (
	"fmt"
	"math"
)

func count_digits(n int64) int64 {
	var res int64
	if n == 0 {
		return 1
	}
	for n != 0 {
		n = n / 10
		res++
	}
	return res
}



func main() {
	var n int64
	fmt.Println("Enter a number")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
	fmt.Printf("number of digits in %d is %d\n", n, count_digits(n))
	fmt.Println("using log: ", int(math.Floor(math.Log10(float64(n)))) + 1)
}