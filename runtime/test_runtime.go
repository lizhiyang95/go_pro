package main

import (
	"fmt"
	"runtime"
	"sync"
)

func testGoexit() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

func testGosched() {
	runtime.GOMAXPROCS(1) // 设置使用一个 CPU 核心
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine 1:", i)
			runtime.Gosched() // 让出当前 goroutine 的执行权
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine 2:", i)
			runtime.Gosched() // 让出当前 goroutine 的执行权
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main goroutine exit")
}

func main() {
	testGosched()
}
