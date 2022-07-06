package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	buf := new(unix.Statfs_t)
	fmt.Printf("Normal mountpoint with error: %v\n", unix.Statfs("/proc", buf))
	fmt.Printf("Mountpoint not exist with error: %v\n", unix.Statfs("/dev/xxxxx", buf))
}
