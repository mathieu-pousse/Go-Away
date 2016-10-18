/**
 * Filename: goroutine.go
 * Author: Nyah Check
 * Usage: Experiments on Goroutines. Adding Mutex's to guarantee channeling
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	lock    sync.Mutex
)

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

	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}

func process() {

	fmt.Println("Processing.")
}
