package main

import "fmt"

func createMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func main() {
	n := 3
	matrix := createMatrix(n)
	for _, row := range matrix {
		fmt.Println(row)
	}
}