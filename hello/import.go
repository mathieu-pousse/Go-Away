/**
 * Filename: simple.go
 * Author: Nyah Check
 * Usage: Simple output
 * Licence: GNU PL 2016
 */
package main


import (

	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	
	fmt.Println("It's over", os.Args[1])
}

