/**
 * Filename: struct.go
 * Author: Nyah Check
 * Usage: Experiments on stuctures.
 * Licence: GNU PL 2016
 */
package main

import (

	"fmt"
)


type p_data struct {
	
		f_name string
		l_name string
		p_age int
		p_temp int	
	}


func main() {

	age, temp := 25, 38
	my_data := p_data{"Abu", "Fayed", age, temp}
	
	print_data(my_data)
}


func print_data(s p_data) int {

	fmt.Println("Here are the Person's details: ")
	fmt.Println("Names ", s.f_name, s.l_name)
	fmt.Println("Age and Body temperature", s.p_age, s.p_temp)
	
	return 0
}
