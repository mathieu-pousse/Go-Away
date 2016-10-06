/**
 * Filename: printfile.go
 * Author: Nyah Check
 * Usage: Prints the contents of a file.
 * Licence: GNU PL 2016
 */
package main


import (

	"fmt"
	"os"
	"io"
)


func main() {

	if len(os.Args) != 2 {
		fmt.Println("No enough command line arguments.")
		os.Exit(1)
	}
	
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Invalid file.", err)
		return
	}
	
	defer file.Close()
	
	input, err := fmt.Scan(file)
	fmt.Printf("File Input\n\n")
	
	for err != io.EOF {
		fmt.Println(string(input))
		input, err = fmt.Scan(file)
	}
}
