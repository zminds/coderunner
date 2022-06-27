package main

import (
	"fmt"
	"os"
)

func main() {
	s := []int{1, 2, 3, 4}
	for _, v := range os.Args {
		fmt.Println(v, s[:])
	}
}
