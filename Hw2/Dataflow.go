package hw2

func Dataflow() {
	// Câu lệnh 1: Định nghĩa cho biến 'x'
	x := 1
	// Câu lệnh 2: Định nghĩa cho biến 'y'
	y := 2

	if x > 0 { // Câu lệnh 3: Điều kiện
		// Câu lệnh 4: Định nghĩa mới cho 'y'
		y = 3
	} else {
		// Câu lệnh 5: Định nghĩa mới cho 'x'
		x = 5
	}

	// Câu lệnh 6: Sử dụng 'x' và 'y'
	z := x + y

	// Câu lệnh 7: Định nghĩa mới cho 'x'
	x = 10

	// Câu lệnh 8: Sử dụng 'x'
	println(z) // Sử dụng 'z'
	println(x) // Sử dụng 'x'
}
