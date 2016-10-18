/**
 * Filename: odd_even.go
 * Author: Nyah Check
 * Usage: Experiments on printing odd and even numbers.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("\n%d is even", i)
		} else {
			fmt.Printf("\n%d is odd", i)
		}
	}
}
