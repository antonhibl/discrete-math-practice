package main

import (
	"fmt"
)

func main() {
	y := 9999
	z := 0
	for x := 1000; x <= y; x++ {
		if x%5 == 0 && x%7 == 0 {
			z++
		}
	}
	fmt.Printf("%d", z)
}
