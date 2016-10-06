/**
 * Filename: var.go
 * Author: Nyah Check
 * Usage: Experiments on variables.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
)

func main() {

	var power int
	power = 9000
	fmt.Println("It's over ", power)

	//making use of multiple var declaration
	first_name, last_name := "Nyah", "Check"

	print_name(first_name)
	i := p_name(first_name)
	if i == 0 {
		sum := add(power, i)
		fmt.Println("Sum is: ", sum)
	} else {
		_, i = print_names(first_name, last_name)

	}

}

func print_name(name string) {

	fmt.Println("Name: ", name)
}

func p_name(name string) int {

	first, last := name, "Check"
	fmt.Println("Name: ", first, last)

	return 0
}

func print_names(name, lname string) (string, int) {

	first, last := name, lname

	fmt.Println("Name: ", first, last)

	return first, 0
}

func add(a, b int) int {

	return a + b
}
