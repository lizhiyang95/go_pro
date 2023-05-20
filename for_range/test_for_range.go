package main

import "fmt"

func testMap() {
	m := make(map[string]string)
	m["name"] = "tom"
	m["age"] = "20"
	for k, v := range m {
		fmt.Printf("%v %v\n", k, v)
	}
}

func testString() {
	s := "abcd"
	for _, v := range s {
		fmt.Printf("v: %s\n", v)
	}
}

func testSlice() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

func testChan() {
	ch := make(chan int)

	// 向管道中写入数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 从管道中读取数据
	for x := range ch {
		fmt.Println(x)
	}
	fmt.Println("done")
}

func main() {
	// testMap()
	// testSlice()
	// testString()
	testChan()
}
