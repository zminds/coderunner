package main

import "fmt"

func main() {
	test := NewTest("hello world")
	fmt.Println(test.name)
	fmt.Println(test.string())
}

func myFunc(param string) error {
	fmt.Println("myFunc: ", param)
	return nil
}

type Test struct {
	name string
}

func (t *Test) string() string {
	return t.name
}

func NewTest(param string) (test *Test) {
	return &Test{
		name: param,
	}
}
