package main

import (
	"fmt"
)

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return (n * fact(n-1))
}

func main() {
	var n, r int
	var result_num, result_denom int

	fmt.Printf("Enter N: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Enter R: ")
	fmt.Scanf("%d", &r)

	result_num = fact(n)
	result_denom = fact(n - r)

	var result int = result_num / result_denom

	fmt.Printf("P(n, r) = %d", result)
}
