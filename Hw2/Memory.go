package main

import "fmt"

type MyClass struct {
	Value int
}

func Memory() {
	// Cấp phát bộ nhớ trên heap
	obj := &MyClass{}
	obj.Value = 42

	fmt.Println(obj.Value)

	// Giải phóng: Go GC tự làm, không có delete
	obj = nil
}
