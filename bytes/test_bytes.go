package main

import (
	"bytes"
	"fmt"
)

// 具体来说，bytes.Runes(b) 函数会遍历字节数组 b 中的每个 UTF-8 编码的字符，并将其转换为 Unicode 编码的整数值，最终返回一个 []rune 切片。
func testRunes() {
	b := []byte("你好世界")
	r := bytes.Runes(b)
	i := len(r)
	fmt.Printf("r: %v\n", r)
	fmt.Printf("i: %v\n", i)
}

func main() {
	testRunes()
}
