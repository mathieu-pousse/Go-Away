/**
 * Filename: fizzBuzz.go
 * Author: Nyah Check
 * Usage: Prints numbers between 1 and 100. But for	multiples	of	three(Fizz)
 * instead	of	the	number,	and	for	the	multiples	of	five(Buzz). For	numbers	that	are	multiples	of
 * both	three	and	five,	print	“FizzBuzz.” that are evenly divisible by three.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("\nFizzBuzz")
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Printf("\nFizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Printf("\nBuzz")
		} else {
			fmt.Printf("\n%d", i)
		}
	}
}
