package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	buf := new(unix.Statfs_t)
	fmt.Printf("%v", unix.Statfs("/dev/sda", buf))
}
