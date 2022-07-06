package main

import (
	"os"
)

func main() {
  f, _ := os.Create("/tmp/sparse.dat")
  f.Seek(1024*1024*10, 0)
  f.Write([]byte("end"))
}