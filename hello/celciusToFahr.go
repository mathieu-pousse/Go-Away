/**
 * Filename: celciusToFahr.go
 * Author: Nyah Check
 * Usage: Converts temperate from celsius to Fahrenheit
 * Licence: GNU GPL 2016
 */
package main

import (
	"fmt"
)

var (
	celsius = 0.0
	fahr    = 0.0
)

//Converts celcius to Fahrenheit using the formula
// F = 32 + 9*C/5
func main() {
	fmt.Println("This program converts the temperature from Celsius to Fahrenheit.")
	fmt.Printf("Enter Celsius temperature: ")
	fmt.Scanf("%f", &celsius)

	fahr = 32 + 9*celsius/5

	fmt.Printf("Celsius temperature %.3f in Fahrenheit is: %.3f", celsius, fahr)

}
