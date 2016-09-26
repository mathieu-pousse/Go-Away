/**
 * Filename: struct.go
 * Author: Nyah Check
 * Usage: Experiments on Arrays, slices and Maps.
 * Licence: GNU GPL 2016
 */
package main

import (
	"fmt"
)

type st_info struct {

	Name string
	ID string
	Age int
	Dept string
}

var students [10]st_info

students[0] = {"Nyah Check", "FE12A025", 25, "Computer Engineering"}
scores := [4]int{78, 84, 56, 69} //students math scores.

func main() {

	for index, value := range scores {
		
	}

}

