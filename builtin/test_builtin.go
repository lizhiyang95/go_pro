package main

import (
	"fmt"
)

// len：返回字符串、数组、切片、映射、通道等的长度。
func testLen() {
	s := "hello, world"
	a := [3]int{1, 2, 3}
	sl := []int{4, 5, 6}
	m := map[string]int{"foo": 1, "bar": 2}
	c := make(chan int, 10)
	fmt.Println(len(s))  // 输出：12
	fmt.Println(len(a))  // 输出：3
	fmt.Println(len(sl)) // 输出：3
	fmt.Println(len(m))  // 输出：2
	fmt.Println(len(c))  // 输出：0
}

// cap：返回切片、数组、通道等的容量。
func testCap() {
	sl := []int{1, 2, 3}
	c := make(chan int, 10)
	fmt.Println(cap(sl)) // 输出：3
	fmt.Println(cap(c))  // 输出：10
}

// make：用于创建切片、映射、通道等数据结构。
func testMake() {
	sl := make([]int, 3, 10)
	m := make(map[string]int)
	c := make(chan int, 10)
	fmt.Printf("sl: %v\n", sl)
	fmt.Printf("m: %v\n", m)
	fmt.Printf("c: %v\n", c)
}

// new：用于创建值类型的指针。
func testNew() {
	i := new(int)
	fmt.Println(*i) // 输出：0
}

// append：用于向切片中追加元素。
func testAddpen() {
	sl := []int{1, 2, 3}
	sl = append(sl, 4, 5, 6)
	fmt.Println(sl) // 输出：[1 2 3 4 5 6]
}

// copy：用于将一个切片的元素复制到另一个切片。
func testCopy() {
	src := []int{1, 2, 3}
	dst := make([]int, 3)
	copy(dst, src)
	fmt.Println(dst) // 输出：[1 2 3]
}

// panic 和 recover：用于实现错误处理和异常处理。
func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	panic("something wrong")
}

func main() {
}
