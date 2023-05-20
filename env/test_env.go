package main

import (
	"fmt"
	"os"
)

func testGetEnv() {
	s := os.Getenv("GOPATH")
	s1 := os.Getenv("GOROOT")
	s2 := os.Environ()
	s3 := os.Getenv("windir")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
}

func testSetEnv() {
	os.Setenv("NAME", "tom")
	s := os.ExpandEnv(" $NAME is me")
	s1 := os.Getenv("NAME")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
}

func testLookupEnv() {
	s1, b1 := os.LookupEnv("GOPATH")
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("b1: %v\n", b1)

	s2, b2 := os.LookupEnv("Path")
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("b2: %v\n", b2)
}

func main() {
	testGetEnv()
	testLookupEnv()
	testSetEnv()
}
