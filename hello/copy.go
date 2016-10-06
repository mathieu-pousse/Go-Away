/**
 * Filename: copy.go
 * Author: Nyah Check
 * Usage: copies values between indexes on slices.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	num := make([]int, 100)

	for i := 0; i < 100; i++ {
		num[i] = int(rand.Int31n(1000))
	}

	sort.Ints(num)

	worst := make([]int, 5)
	copy(worst, num[:len(num)])

	fmt.Println(worst)
	fmt.Printf("\nOriginal Array")
	//fmt.Println(num);

	//print the array in reverse
	fmt.Printf("\n\nPrints the Reverse Array.")
	//for i := len(num) - 1; i >= 0; i-- {
	//	println(num[i])
	//}
}
