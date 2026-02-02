package main

import "fmt"

func mySqrt_recursive(n int) int {
	if n < 2 {
		return n
	}
	target := int64(n)
	var low int64 = 2
	var high int64 = target/2
	var ans int64 = 1
	for low <= high {
		mid := low + (high-low)/2
		midsq := mid * mid
		if midsq == target {
			return int(mid)
		} else if midsq < target {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return int(ans)
}


func main() {
	testCases := []int{0, 1, 8, 16, 2147483647}

	for _, x := range testCases {
		fmt.Printf("sqrt(%d) rounded down = %d\n", x, mySqrt_recursive(x))
	}
}