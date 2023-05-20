package main

import "fmt"

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type Sayer interface {
	say()
}

type FlyFish interface {
	Flyer
	Swimmer
	Sayer
}

type Fish struct {
}

func (this Fish) fly() {
	fmt.Println("fly")
}

func (this Fish) swim() {
	fmt.Println("swim")
}

func (this Fish) say() {
	fmt.Println("say")
}

func (this Fish) polymorphic(fish FlyFish) {
	fish.fly()
	fish.say()
	fish.swim()
}

func NewFish() Fish {
	return Fish{}
}

func main() {
	f := NewFish()
	f.fly()
	f.say()
	f.swim()

	f.polymorphic(f)
}
