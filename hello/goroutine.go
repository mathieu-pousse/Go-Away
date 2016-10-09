/**
 * Filename: goroutine.go
 * Author: Nyah Check
 * Usage: Experiments on Goroutines.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
	"time"
)

var counter = 0

func main() {

	fmt.Println("Start")

	for i := 0; i < 20; i++ {
		go incr()
	}

	//go process() //starts goroutine.

	time.Sleep(time.Millisecond * 10)
	fmt.Println("Done!")
}

func incr() {

	counter++
	fmt.Println(counter)
}

func process() {

	fmt.Println("Processing.")
}
