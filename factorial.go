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
	var n int

	fmt.Printf("Enter N: ")
	fmt.Scanf("%d", &n)

	var x int = fact(n)

	fmt.Printf("n! = %d", x)
}
