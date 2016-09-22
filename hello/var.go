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
	
	//making user of multiple var declaration
	first_name, last_name := "Nyah", "Check"
	
	fmt.Println("His name is: ", first_name, last_name)

}
