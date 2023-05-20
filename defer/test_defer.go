package main

import (
	"fmt"
)

func testDeferIndex() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
}

// 捕获的是祖父级（相等）调用时的异常，直接调用时无效
func testPanicInvalid() {
	recover()
	panic("无效情况")
}

func testPanicInvalid2() {
	defer recover()
	panic("无效情况")
}

func testPanicInvalid3() {
	defer func() {
		func() {
			recover()
		}()
	}()
	panic("无效情况")
}
func testPanicValid() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()
	panic("有效情况")
}

func main() {
	testDeferIndex()
	// testPanicInvalid()
	// testPanicInvalid2()
	// testPanicInvalid3()
	testPanicValid()
}
