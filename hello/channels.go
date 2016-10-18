/**
 * Filename: channels.go
 * Author: Nyah Check
 * Usage: Uses channels to synchronize data between processes.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

//Worker function for channeling data.
func (w Worker) process(c chan int) {

	data := <-c
	fmt.Printf("\nWorker %d, received %d data.", w.id, data)
}

func main() {

	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c) //passes data to channel for processing

	}

	for {
		select {
		case c <- rand.Int():
			time.Sleep(time.Millisecond * 50)

		case <-time.After(time.Millisecond * 100):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}
	}
}
