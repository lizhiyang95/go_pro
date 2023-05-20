package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var reader strings.Reader

func createReader() {
	r := strings.NewReader("hello golang")
	b := make([]byte, 20)
	n, _ := r.Read(b)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("b: %v\n", string(b))
}

func copy() {
	r := strings.NewReader("hello java")

	/* 	//拷贝到文件
	   	f, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	   	_, err := io.Copy(f, r) */
	//拷贝到标准输出
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		log.Fatal(err)
	}
}

func copyByBuffer() {
	r := strings.NewReader("first buffer\n")
	r2 := strings.NewReader("second buffer\n")
	var b []byte

	fmt.Printf("b: %v\n", b)
	_, err := io.CopyBuffer(os.Stdout, r, b)
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := io.CopyBuffer(os.Stdout, r2, b)
	if err2 != nil {
		log.Fatal(err2)
	}

}

func main() {
	//createReader()
	//copy()
	copyByBuffer()
}
