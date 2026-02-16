package main

import "fmt"

func modularAddition(a, b, m int64) int64 {
    // To prevent overflow: (a % m + b % m) % m
    return ((a % m) + (b % m)) % m
}

func main() {
    var a, b, m int64 = 1000000000, 1000000000, 7
    fmt.Println(modularAddition(a, b, m))
}