package main

import (
	"fmt"
	"log"
	"os"
)

func testLog() {
	log.Print("log")
	log.Printf("log %v ", 100)
	log.Println("my", "log")
}

func testPanic() {
	defer fmt.Println("defer")
	log.Panic("panic")
}

func testFatal() {
	defer log.Print("defer")
	log.Fatal("fatal")
}

func logFile() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetPrefix("mylog:")
	f, err := os.OpenFile("log/a.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal("文件打开错误")
	}
	log.SetOutput(f)
}

func init() {
	logFile()
}

func main() {
	testFatal()
	testPanic()
}
