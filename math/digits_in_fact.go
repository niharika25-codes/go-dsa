package main

import (
	"fmt"
	"math"
)


func countDigitsInFactorial(n int) int {
    if n < 0 { return 0 }
    if n <= 1 { return 1 }

    var logSum float64
    for i := 1; i <= n; i++ {
        logSum += math.Log10(float64(i))
    }

    return int(math.Floor(logSum)) + 1
}

func main() {
	var n float64
	n = 1234
	if n < 0 {
		fmt.Println(0)
	}
	if n <= 1 {
		fmt.Println(1)
	}

	x := (n * math.Log10(n/math.E) + math.Log10(2*math.Pi*n)/2)
	fmt.Println(math.Floor(x)+1)
}