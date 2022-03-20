package main

import "fmt"

func main() {
	fmt.Println(NewTest("hello world").name)
	fmt.Println(NewTest("hello world").string())
}

func myFunc(param string) error {
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
