package main

// import "fmt"

// //	func main() {
// //		fmt.Println(os.Getenv("APP_NAME"))
// //	}
// //

// //	func add(x int, y int) int {
// //		return x + y
// //	}
// // func add(x, y int) int {
// // 	return x + y
// // }
// // func main() {
// // 	fmt.Print(add(12, 12))
// // }

// // func swap(x, y string) (string, string) {
// // 	return y, x
// // }
// // func main() {
// // 	a, b := swap("hello", "world")
// // 	fmt.Println(a, b)
// // }

// // func split(sum int) (x, y int) {
// // 	x = sum * 4 / 9
// // 	y = sum - x
// // 	return
// // }
// // func main() {
// // 	fmt.Println(split(17))
// // }

// // Khai báo biến
// // var i, j int = 1, 2

// // func main() {
// // 	var c, python, java = true, false, "no!"
// // 	k := 3
// // 	fmt.Println(i, j, c, python, java, k)
// // }

// // End Khai báo biến

// //  Vòng lặp for

// // func main() {
// // 	sum := 0
// // 	for i := 0; i < 10; i++ {
// // 		sum += i
// // 	}
// // 	fmt.Println(sum)
// // }

// // func main() {
// // 	sum := 1
// // 	for sum < 1000 {
// // 		sum += sum
// // 	}
// // 	fmt.Println(sum)
// // }

// // Vòng for vô tận
// // for {}

// //  End Vòng lặp for

// // If
// // func pow(x, n, lim float64) float64 {
// // 	if v := math.Pow(x, n); v < lim {
// // 		return v
// // 	}
// // 	return lim
// // }

// // func main() {
// // 	fmt.Println(
// // 		pow(3, 2, 10),
// // 		pow(3, 2, 20),
// // 	)
// // }
// // End If

// // Defer : thực thi ngay khi hàm kết thúc

// // Defer : có tính tracking
// // func main() {
// // 	fmt.Println("counting")
// // 	for i := 0; i < 10; i++ {
// // 		defer fmt.Println(i)
// // 	}
// // 	defer fmt.Println("done")
// // }

// // Pointer
// // func main() {
// // 	i, j := 42, 2701
// // 	p := &i
// // 	fmt.Println(*p)
// // 	*p = 21
// // 	fmt.Println(i)

// // 	p = &j
// // 	*p = *p / 37
// // 	fmt.Println(j)
// // }

// // Struct
// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	// fmt.Println(v)
// 	p := &v
// 	p.X = 1e9
// 	fmt.Println(v)
// }
