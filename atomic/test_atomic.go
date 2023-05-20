package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func load() {
	i2 := atomic.LoadInt32(&i)
	fmt.Printf("i2: %v\n", i2)
}

func store() {
	atomic.StoreInt32(&i, 200)
	fmt.Printf("i: %v\n", i)
}

func compare() {
	b := atomic.CompareAndSwapInt32(&i, 100, 400)
	fmt.Printf("b: %v\n", b)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	load()
	store()
	compare()

	//这里等待一秒，不然输出时还没运行完
	time.Sleep(time.Second)
	fmt.Printf("i: %v\n", i)
}
