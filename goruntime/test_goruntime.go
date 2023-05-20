package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 无缓冲，接收和发送会阻塞等待，直到报错，可以采用go协程解决
var channel = make(chan int)

// 有缓冲，先写入缓冲，缓冲满了和上面一样
//var channel = make(chan int, 1)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100)
	fmt.Printf("value1: %v\n", value)
	//time.Sleep(time.Second * 2)
	channel <- value
}

func receive() {
	defer close(channel)
	value := <-channel
	fmt.Printf("receive : %v\n", value)
}

func main() {
	go send()
	fmt.Println("wait...")
	receive()
	fmt.Println("end...")

	go func() {
		time.Sleep(time.Second)
		fmt.Println("wait...")
	}()
	runtime.Gosched()
	fmt.Println("end...")
}
