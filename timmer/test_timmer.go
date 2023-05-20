package main

import (
	"fmt"
	"time"
)

func timmer() {
	t := time.NewTimer(time.Second * 2)
	//fmt.Printf("t: %v\n", t)
	fmt.Printf("time.Now(): %v\n", time.Now())
	//阻塞2秒
	c := <-t.C
	fmt.Printf("c: %v\n", c)
}

func after() {
	<-time.After(time.Second * 2)
	fmt.Println("after...")
}

func chanBlock() {
	var a = make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		a <- 123
	}()

	<-a
}

func main() {
	timmer()
	//after()

	//chanBlock()
	fmt.Println("end....")
}
