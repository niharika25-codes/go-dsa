package main

import "fmt"

func reverse(a []int) {
	for s, e := 0, len(a)-1; s < e; s, e = s+1, e-1 {
		a[s], a[e] = a[e], a[s]
	}
}

func main() {
	rev := []int{1,2,3,4,5}
	 reverse(rev)
	fmt.Println("rev: ", rev)
}