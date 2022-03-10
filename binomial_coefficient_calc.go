package main

import (
	"fmt"
)

func main() {
	var k, n, nth int

	fmt.Printf("Enter N: ")
	fmt.Scanln(&n)
	fmt.Printf("Enter kth term: ")
	fmt.Scanln(&nth)

	for k = 0; k <= n; k++ {
		if k+1 == nth {
			fmt.Printf("(%d!/(%d!(%d-%d)!))x^%d*y^%d", n-k, k, n-k, k, n-k, k)
		}
	}
}
