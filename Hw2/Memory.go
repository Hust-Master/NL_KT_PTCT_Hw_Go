package hw2

import "fmt"

type Counter struct {
	value int
}

func Memory() {
	fmt.Println("===================== 2. Memory =====================")

	// --- A. Cấp phát Bộ nhớ Tĩnh (Stack) ---
	// Biến cục bộ 'count1' được cấp phát trên Stack (thường)
	// Vòng đời của nó kết thúc khi hàm main() kết thúc.
	var count1 int = 10
	fmt.Printf("1. Biến Stack (int): Địa chỉ = %p, Giá trị = %d\n", &count1, count1)

	// Cấu trúc 'c1' cũng được cấp phát trên Stack (thường)
	c1 := Counter{value: 50}
	fmt.Printf("2. Struct Stack: Địa chỉ = %p, Giá trị = %d\n", &c1, c1.value)

	fmt.Println("---")

	// --- B. Cấp phát Bộ nhớ Động (Heap) ---
	// Sử dụng 'new' để cấp phát động. Biến 'count2Ptr' là con trỏ,
	// còn dữ liệu thực tế (*count2Ptr) được cấp phát trên Heap.
	count2Ptr := new(int) // Cấp phát 1 int trên Heap và trả về con trỏ
	*count2Ptr = 20
	fmt.Printf("3. Biến Heap (int): Con trỏ = %p, Địa chỉ giá trị = %p, Giá trị = %d\n", &count2Ptr, count2Ptr, *count2Ptr)

	// Cấu trúc 'c2' được cấp phát trên Heap (thường, do dùng 'new')
	c2Ptr := new(Counter)
	c2Ptr.value = 60
	fmt.Printf("4. Struct Heap: Con trỏ = %p, Địa chỉ giá trị = %p, Giá trị = %d\n", &c2Ptr, c2Ptr, c2Ptr.value)

	// Biến 'count3' có thể bị Go "escape analysis" (phân tích thoát)
	// đẩy lên Heap nếu nó được truyền ra ngoài hàm main
	count3 := escapeToHeap(&c1) // Truyền địa chỉ của c1 để minh họa
	fmt.Printf("5. Biến sau Escape (Struct): Địa chỉ = %p, Giá trị = %d (Trong TH c1 bị đẩy lên Heap)\n", count3, count3.value)

	fmt.Println("---")

	// --- C. Quản lý Con trỏ (Pointer) và Giải phóng Bộ nhớ ---
	// Go không sử dụng 'free()' mà dùng Garbage Collector (GC).
	// Khi một đối tượng trên Heap không còn con trỏ nào trỏ tới, GC sẽ tự động giải phóng nó.

	// Tạo một đối tượng trên Heap
	dataPtr := new(string)
	*dataPtr = "Dữ liệu quan trọng"
	fmt.Printf("6. Trước giải phóng: Con trỏ Data = %p, 'Giá trị dataPtr' = %s\n", dataPtr, *dataPtr)

	// Để "giải phóng" (hay làm cho nó đủ điều kiện để GC thu gom):
	dataPtr = nil // Đặt con trỏ về nil. Đối tượng Heap không còn tham chiếu nào.

	fmt.Printf("7. Sau 'Giá trị dataPtr = nil': Con trỏ Data = %p\n", dataPtr)
	// Sau khi gán nil, nếu không có con trỏ nào khác trỏ tới vùng nhớ đó,
	// Garbage Collector của Go sẽ tự động giải phóng vùng bộ nhớ Heap sau đó.

	// Hàm này (hoặc bất kỳ hàm nào khác) không cần gọi 'free(dataPtr)'.
	fmt.Println("=================================")
}

// Hàm này có thể khiến tham số 'c' bị cấp phát trên Heap (Escape Analysis)
func escapeToHeap(c *Counter) *Counter {
	// Nếu không có return, c sẽ ở Stack (vì c1Ptr không thoát ra ngoài)
	// Nhưng vì nó được trả về, c1Ptr (con trỏ tới c1) "thoát" ra khỏi hàm.
	return c
}
