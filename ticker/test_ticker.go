package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan int, 0)

var ticker = time.NewTicker(time.Second)

func tick() {
	counter := 0
	for v := range ticker.C {
		fmt.Printf("v: %v\n", v)
		counter++
		if counter >= 5 {
			ticker.Stop()
			break
		}
	}
}

func tickSelect() {
	for _ = range ticker.C {
		select {
		case chan1 <- 1:
		case chan1 <- 2:
		case chan1 <- 3:
		}
	}
}

func tickReceive() {
	sum := 0
	for v := range chan1 {
		fmt.Printf("v: %v\n", v)
		sum += v
		if sum >= 15 {
			break
		}
	}
}

func main() {
	//tick()
	go tickSelect()
	tickReceive()
}
