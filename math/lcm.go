package main

import "fmt"

func lcm(a, b int) int {
	res := max(a, b)
	for {
		if res%a == 0 && res%b == 0 {
			return res
		}
		res++
	}
	return res
}

func main() {
	fmt.Println("lcm(12, 15): ", lcm(12, 15))
}