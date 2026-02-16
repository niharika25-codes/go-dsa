package main

import "fmt"

// findTwoOddOccurrences finds two numbers that appear an odd number of times in an array.
// All other numbers appear an even number of times.
func findTwoOddOccurrences(arr []int) (int, int) {
	// Step 1: XOR all elements to get the XOR sum of the two odd numbers.
	// xorSum will be num1 ^ num2
	xorSum := 0
	for _, num := range arr {
		xorSum ^= num
	}

	// Step 2: Find a distinguishing bit (rightmost set bit) in xorSum.
	// This bit will be set in one of the odd numbers and unset in the other.
	// The expression `xorSum & (-xorSum)` isolates the rightmost set bit.
	rightmostSetBit := xorSum & (-xorSum)

	// Step 3 & 4: Divide numbers into two groups and XOR within each group.
	num1 := 0
	num2 := 0
	for _, num := range arr {
		if (num & rightmostSetBit) != 0 {
			// Number belongs to the first group
			num1 ^= num
		} else {
			// Number belongs to the second group
			num2 ^= num
		}
	}

	return num1, num2
}

func main() {
	arr1 := []int{5, 5, 2, 8, 1, 8, 1, 1}
	// Expected output: (2, 1) or (1, 2)
	result1_num1, result1_num2 := findTwoOddOccurrences(arr1)
	fmt.Printf("Array: %v, Two odd occurring numbers: (%d, %d)\n", arr1, result1_num1, result1_num2)

	arr2 := []int{10, 7, 6, 6, 6, 6, 7, 4}
	// Expected output: (10, 4) or (4, 10)
	result2_num1, result2_num2 := findTwoOddOccurrences(arr2)
	fmt.Printf("Array: %v, Two odd occurring numbers: (%d, %d)\n", arr2, result2_num1, result2_num2)

	arr3 := []int{12, 23, 34, 12, 12, 23, 12, 45}
	// Expected output: (34, 45) or (45, 34)
	result3_num1, result3_num2 := findTwoOddOccurrences(arr3)
	fmt.Printf("Array: %v, Two odd occurring numbers: (%d, %d)\n", arr3, result3_num1, result3_num2)
}
