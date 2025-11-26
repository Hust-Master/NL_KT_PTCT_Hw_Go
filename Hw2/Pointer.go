package hw2

import (
	"fmt"
	"runtime"
)

type Data struct {
	Value int
	Name  string
}

func Pointer() {
	fmt.Println("===================== 3. Point =====================")
	fmt.Println("=== Khởi tạo và theo dõi bộ nhớ (trước cấp phát) ===")
	printMemStats()

	// 1. Cấp phát bộ nhớ động cho một giá trị (Sử dụng con trỏ)
	// 'new(int)' cấp phát vùng nhớ cho một số nguyên (int) và trả về con trỏ *int
	ptrA := new(int)
	*ptrA = 100                                                      // Gán giá trị 100 vào vùng nhớ mà ptrA trỏ tới
	fmt.Printf("\nĐịa chỉ của ptrA: %p, Giá trị: %d\n", ptrA, *ptrA) // %p in địa chỉ của con trỏ

	// 2. Cấp phát bộ nhớ động cho một cấu trúc (struct)
	// 'ptrB' là một con trỏ (*Data) trỏ đến một cấu trúc Data mới trên Heap.
	ptrB := &Data{Value: 200, Name: "Dynamic Data"}
	fmt.Printf("Địa chỉ của ptrB: %p, Giá trị: %+v\n", ptrB, *ptrB)

	// 3. Kết thúc phạm vi sử dụng (Scope)
	// Sau khi hàm main() kết thúc, các con trỏ ptrA và ptrB không còn được sử dụng.
	// Dù không có lệnh 'free()' thủ công, Garbage Collector sẽ tự động
	// thu hồi bộ nhớ mà chúng đang trỏ tới.

	fmt.Println("\n=== Kích hoạt Garbage Collector (bộ thu gom rác) ===")
	// Gọi thủ công để minh họa việc giải phóng, dù thường nó xảy ra tự động.
	runtime.GC()
	fmt.Println("=== Theo dõi bộ nhớ (sau Garbage Collector) ===")
	printMemStats()
	fmt.Println("==========================================")
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// HeapAlloc là tổng bộ nhớ heap được cấp phát (bytes)
	fmt.Printf("Bộ nhớ Heap đang sử dụng: %v Bytes\n", m.HeapAlloc)
}
