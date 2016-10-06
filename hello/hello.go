/**
 * Filename: Hello.go
 * Author: Nyah Check
 * Usage: Hello world in Golang
 * Licence: GNU PL 2016
 * Buea, Cameroon.
 */

package main

import (
	"fmt"
	"github.com/Ch3ck/hello/stringutil"
)

func main() {

	fmt.Printf(stringutil.Reverse("!oG, olleH"))
	println("It's currently 9.59am in the morning!")
}
