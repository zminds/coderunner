package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("FEATURE_ENCRYPTION_KEY", "abcdefghijklmnop")
	ba := []byte(os.Getenv("FEATURE_ENCRYPTION_KEY"))
	fmt.Println(ba)
	fmt.Println(len(ba))
}
