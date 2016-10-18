/**
 * Filename: numToWords.go
 * Author: Nyah Check
 * Usage: Experiments converting integers to words.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		switch i {
		case 1:
			fmt.Printf("\nOne")
		case 2:
			fmt.Printf("\nTwo")
		case 3:
			fmt.Printf("\nThree")
		case 4:
			fmt.Printf("\nFour")
		case 5:
			fmt.Printf("\nFive")
		case 6:
			fmt.Printf("\nSix")
		case 7:
			fmt.Printf("\nSeven")
		case 8:
			fmt.Printf("\nEight")
		case 9:
			fmt.Printf("\nNine")
		default:
			fmt.Printf("\nZero")
		}
	}
}
