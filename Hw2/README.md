# Homework 2
Lựa chọn 1 công cụ/ngôn ngữ lập trình mà nhóm thống nhất là nắm vững nhất:
1. Minh họa ý tưởng phân tích chương trình sử dụng Dataflow Analysis
2. Minh họa ý tưởng phân tích mã nguồn chương trình sử dụng kỹ thuật CFG
3. Minh họa các thuật cấp phát bộ nhớ và giải phóng bộ nhớ
4. Minh họa ý tưởng sử dụng và giải phóng pointer


## 1. Minh họa ý tưởng phân tích chương trình sử dụng Dataflow Analysis
Dataflow Analysis là kỹ thuật theo dõi giá trị biến tại từng điểm trong chương trình. Ví dụ: xác định giá trị có thể của biến x tại mỗi dòng.

```go
package main

import "fmt"

func main() {
    x := 0   // Dòng 1
    y := 5   // Dòng 2

    if y > 3 { // Dòng 3
        x = y + 1 // Dòng 4
    } else {
        x = y - 1 // Dòng 5
    }

    fmt.Println(x) // Dòng 6
}

```

Phân tích:
- Trước dòng 3: x=0, y=5
- Nếu nhánh if đi vào (đúng): x = 6
- Nếu nhánh else đi vào (sai): x = 4
- Dòng 6: x ∈ {4,6}  

Đây là lấy tập giá trị khả dĩ cho mỗi biến tại mỗi điểm, ý tưởng cơ bản của Dataflow Analysis.


## 2. Minh họa ý tưởng phân tích mã nguồn chương trình sử dụng kỹ thuật CFG


## 3. Minh họa các thuật cấp phát bộ nhớ và giải phóng bộ nhớ
Trong Go, cấp phát bộ nhớ bằng new() hoặc &struct{}. Go có garbage collector, không cần free thủ công.
```go
package main

import "fmt"

type MyClass struct {
    Value int
}

func main() {
    // Cấp phát bộ nhớ trên heap
    obj := &MyClass{}
    obj.Value = 42

    fmt.Println(obj.Value)

    // Giải phóng: Go GC tự làm, không có delete
    obj = nil
}

```


## 4. Minh họa ý tưởng sử dụng và giải phóng pointer

```go
package main

import "fmt"

func main() {
    x := 10
    px := &x // Lấy địa chỉ của x

    fmt.Println("Value via pointer:", *px)

    *px = 20 // Gán giá trị qua pointer
    fmt.Println("New value:", x)
}

```

Chú ý:
- Go không cho pointer arithmetic
- Stack/heap được Go tự quyết