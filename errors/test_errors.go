package main

import (
	"errors"
	"fmt"
	"time"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		//s = err.Error()
		return "", err
	} else {
		return s, nil
	}
}

type Myerror struct {
	When time.Time
	What string
}

func (this Myerror) Error() string {
	return fmt.Sprintf("%v %v", this.When, this.What)
}
func oops() error {
	return Myerror{
		time.Now(),
		"the file system has gone away",
	}
}

func testOops() {
	err := oops()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("err: %v\n", err.Error())
	}
}

func main() {
	testOops()
}
