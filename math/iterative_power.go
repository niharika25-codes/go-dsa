package main

import "fmt"

func power(x, y int) int {
	res := 1
	for y > 0 {
		if y%2 != 0 {
			res *= x
		}
		x *= x
		y /= 2
	}
	return res
}

func main() {
	fmt.Println(power(7, 9))
}