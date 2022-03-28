package main

import (
	"errors"
	"fmt"
)

func main() {
	// test := NewTest("hello world")
	// fmt.Println(test.name)
	// fmt.Println(test.string())
	fmt.Print("test")
}

func myFunc(param string) error {
	fmt.Println("myFunc: ", param)
	if param == "5678" {
		return errors.New("test")
	}
	return nil
}
