package main

import "fmt"

type Animol struct {
	name string
}

func (this *Animol) eat() {
	fmt.Printf("this.name: %v\n", this.name)
	fmt.Printf("\"动物吃东西\": %v\n", "动物吃东西")
}

type Dog struct {
	Animol
}

func (this *Dog) eat() {
	fmt.Printf("this.name: %v\n", this.name)
	fmt.Printf("\"狗吃骨头\": %v\n", "狗吃骨头")
}

func NewDog() *Dog {
	d := &Dog{
		Animol: Animol{
			name: "小黑",
		},
	}
	return d
}

func main() {
	d := NewDog()
	d.eat()
	d.Animol.eat()
}
