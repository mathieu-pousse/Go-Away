/**
 * Filename: slice.go
 * Author: Nyah Check
 * Usage: Experiments on slices expanding with cap and append functions.
 * Licence: GNU PL 2016
 */
package main

import (
	"fmt"
)

func main() {

	scores := make([]int, 0, 5)

	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		//Grow array to accommodate new size
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}

	}

	//tries copying information from one slice to another
	score := []int{1, 2, 3, 4, 5}

	slice := score[2:4]
	slice[0] = 948
	fmt.Println(score)
	fmt.Println(slice)

}
