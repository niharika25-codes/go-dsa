package main

import "fmt"

func power(x, y int) int {
	temp := 0
	if y == 0 {
		return 1
	}
	temp = power(x, (y / 2))
	if y%2 == 0 {
		return temp * temp
	} else {
		return x * temp * temp
	}
}

func main() {
	fmt.Println(power(5, 9))
}