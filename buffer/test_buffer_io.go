package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func testReader() {
	r := strings.NewReader("hello java")
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func testWriter() {
	b := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(b)
	w.WriteString("hello java")
	w.Reset(b)
	// w.Flush()
	fmt.Printf("b: %v\n", b)
}

func main() {
	// testReader()
	testWriter()
}
