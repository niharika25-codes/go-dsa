package main

import (
	"fmt"
	"math"
)

func getSecLarg(nums []int) (int, error) {
	
	if len(nums) < 2 {
		return 0, fmt.Errorf("list must have atleast two elements")
	}
	lar, slar := math.MinInt, math.MinInt
	for _, val := range nums {
        if val > lar {
            slar = lar
            lar = val
        } else if val > slar && val < lar {
            slar = val
        }
    }
    if slar == math.MinInt {
        return 0, fmt.Errorf("no distinct second largest element found")
    }
	return slar, nil
}

func main() {
	nums := []int{34, 55, 9, 1, 23, 50, 60, 79}
	res, err := getSecLarg(nums)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Second Largest: ", res)
    }
}