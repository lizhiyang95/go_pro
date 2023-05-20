package main

import (
	"fmt"
	"time"
)

func testChannel() {
	var chan1 = make(chan int)
	go func() {
		chan1 <- 1
		chan1 <- 2
		chan1 <- 3
		close(chan1)
	}()

	for v := range chan1 {
		fmt.Printf("v: %v\n", v)
	}
}

func testSelect() {

	var chanInt = make(chan int, 0)
	var chanString = make(chan string, 0)
	go func() {
		for {
			chanInt <- 123
			// defer close(chanInt)
			time.Sleep(time.Second)
		}

	}()

	go func() {
		for {
			chanString <- "golang"
			// defer close(chanString)
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		select {
		case r := <-chanInt:

			fmt.Printf("r: %v\n", r)
		case r := <-chanString:
			fmt.Printf("r: %v\n", r)

		case <-time.After(time.Second * 3):
			fmt.Println("after")
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}

//单向通道

func main() {
	testChannel()
	// testSelect()
}
