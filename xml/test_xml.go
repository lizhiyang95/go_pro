package main

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"os"
)

type Person3 struct {
	XMLName xml.Name //xml固定格式
	Name    string
	Age     int
}

func toXml() {
	p := Person3{
		Name: "tom",
		Age:  10,
	}
	// b, _ := xml.Marshal(p)
	b, _ := xml.MarshalIndent(p, " ", "  ")
	s := string(b)
	fmt.Printf("s: %v\n", s)
}

func xmlTo() {
	s := `<Person3>
			<Name>tom</Name>
			<Age>10</Age>
		</Person3>`
	b := []byte(s)
	var p Person3
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func xmlToFile() {
	p := Person3{
		Name: "tom",
		Age:  10,
	}
	f, _ := os.OpenFile("a.xml", os.O_CREATE|os.O_APPEND|os.O_RDWR, fs.ModePerm)
	e := xml.NewEncoder(f)
	e.Encode(p)

}

func fileToXml() {
	f, _ := os.Open("a.xml")
	defer f.Close()
	d := xml.NewDecoder(f)
	var v Person3
	err := d.Decode(&v)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("v: %v\n", v)
}

func main() {
	toXml()
	xmlTo()
	//xmlToFile()
	fileToXml()
}
