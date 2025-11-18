package main

import "fmt"

func Pointer() {
	x := 10
	px := &x // Lấy địa chỉ của x

	fmt.Println("Value via pointer:", *px)

	*px = 20 // Gán giá trị qua pointer
	fmt.Println("New value:", x)
}
