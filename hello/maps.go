/**
 * Filename: Maps.go
 * Author: Nyah Check
 * Usage: Experiments on Maps with Go.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
)

func main() {

	data := make(map[string]int)

	data["goku"] = 9001
	//power, exits := data["vegeta"]

	//gets total number of keys
	//total := len(data)

	//deletes a key from the map
	delete(data, "goku")

	fmt.Println(data)

	//grow map dynamically
	data = map[string]int{

		"Goku":    001,
		"Vegeta":  002,
		"Gohan":   003,
		"Freezer": 004,
		"Piccolo": 005,
		"Goten":   006,
	}

	//iterate and print map
	for value, key := range data {

		fmt.Printf("\nKey: %d, Value: %s", key, value)
	}

}
