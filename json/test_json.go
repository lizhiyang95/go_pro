package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

type Person struct {
	Name string
	Age  int
}

type Person2 struct {
	Name    string
	Age     int
	Money   []int
	Friends map[string]string
}

func toJson() {
	p := Person{
		"tom",
		10,
	}
	b, _ := json.Marshal(p)
	s := string(b)
	fmt.Printf("s: %v\n", s)
}

func jsonTo() {
	b := []byte(`{"Name":"tom","Age":10}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func MulToJson() {
	p := Person2{
		"tom",
		10,
		[]int{1, 2, 3, 4, 5},
		map[string]string{"male": "tom", "female": "kite"},
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func JsonToMul() {
	b := []byte(`{"Name":"tom","Age":10,"Money":[1,2,3,4,5],"Friends":{"female":"kite","male":"tom"}}`)
	var p map[string]any
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
	for _, v := range p {
		fmt.Printf("v: %v\n", v)
	}

}

func fileToJson() {
	f, _ := os.Open("a.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var v map[string]any
	d.Decode(&v)
	fmt.Printf("v: %v\n", v)
}

func jsonToFile() {
	v := Person2{
		"tom",
		10,
		[]int{1, 2, 3, 4, 5},
		map[string]string{"male": "jack", "female": "lucy"},
	}
	f, _ := os.OpenFile("a.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, fs.ModePerm)
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(&v)
	fmt.Printf("v: %v\n", v)
}

func main() {
	toJson()
	//jsonTo()
	//MulToJson()
	//JsonToMul()
	//fileToJson()
	//jsonToFile()
}
