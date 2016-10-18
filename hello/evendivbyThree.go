/**
 * Filename: evendivbyThree.go
 * Author: Nyah Check
 * Usage: Prints numbers between 1 and 100 that are evenly divisible by three.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Printf("\n%d", i)
		}
	}
}
