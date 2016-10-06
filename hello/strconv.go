/**
 * Filename: strconv.go
 * Author: Nyah Check
 * Usage: Converts strings to integers
 * Licence: GNU PL 2016
 */
package main


import (

	"fmt"
	"os"
	"strconv"
)

func main() {


	if len(os.Args) != 2 {
		os.Exit(1)
	}
	
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Not a valid number: ", os.Args[1])
	} else {
		fmt.Println(n)
	}
}
