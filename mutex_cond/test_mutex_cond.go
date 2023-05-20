package main

import (
	"fmt"
	"sync"
	"time"
)

func testMutexCond() {
	var mutex sync.Mutex
	var cond sync.Cond

	// 初始化条件变量
	cond.L = &mutex

	// 协程 A
	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			fmt.Println("协程 A 执行")
			cond.Signal()
			cond.Wait()
			mutex.Unlock()
		}
	}()

	// 协程 B
	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			fmt.Println("协程 B 执行")
			cond.Signal()
			cond.Wait()
			mutex.Unlock()
		}
	}()

	// 等待协程执行完成
	// select {}
	time.Sleep(time.Second * 5)
}

var i = 100

var wg sync.WaitGroup

var lock sync.Mutex

var cond sync.Cond

func add() {
	wg.Done()
	lock.Lock()
	defer lock.Unlock()
	i += 1
	fmt.Printf("i++: %v\n", i)
	cond.Signal()
	cond.Wait()
}

func sub() {
	wg.Done()
	lock.Lock()
	defer lock.Unlock()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	cond.Wait()
	cond.Signal()
}

func testMutexCondWaitgroup() {

	cond.L = &lock

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go sub()
		wg.Add(1)
		go add()
	}

	wg.Wait()
	fmt.Printf("i: %v\n", i)
}

func main() {
	// testMutexCond()
	testMutexCondWaitgroup()
}
