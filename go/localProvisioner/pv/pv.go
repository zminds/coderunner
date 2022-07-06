package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	h := fnv.New32a()
	fmt.Println(h.Sum32())
	h.Write([]byte("mount1"))
	h.Write([]byte("test-node"))
	h.Write([]byte("sc1"))

	fmt.Println(fmt.Sprintf("local-pv-%x", h.Sum32()))
}
