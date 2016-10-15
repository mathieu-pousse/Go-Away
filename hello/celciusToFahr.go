/**
 * Filename: celciusToFahr.go
 * Author: Nyah Check
 * Usage: Converts temperate from celcius to Fahrenheit
 * Licence: GNU GPL 2016
 */
package main

import (
	"fmt"
)

var (
	celcius = 0.0
	fahr    = 0.0
)

//Converts celcius to Fahrenheit using the formula
// F = 32 + 9*C/5
func main() {
	fmt.Println("This program converts the temperature from Celsius to Fahrenheit.")
	fmt.Printf("Enter Celsius temperature: ")
	fmt.Scanf("%f", &celcius)

	fahr = 32 + 9*celcius/5

	fmt.Printf("Celcius temperature %.3f in Fahrenheit is: %.3f", celcius, fahr)

}
