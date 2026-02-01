package main

import "fmt"

func palindrome_num(n int) bool {
	temp := n
	rev := 0
	for temp != 0 {
		ld := temp % 10
		rev = (rev*10) + ld
		temp /= 10
	}
	return rev == n
}

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)
	fmt.Printf("is %d palindrome: %t", n, palindrome_num(n))
}