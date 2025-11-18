package main

import "fmt"

func Dataflow() {
	x := 0 // Dòng 1
	y := 5 // Dòng 2

	if y > 3 { // Dòng 3
		x = y + 1 // Dòng 4
	} else {
		x = y - 1 // Dòng 5
	}

	fmt.Println(x) // Dòng 6
}
