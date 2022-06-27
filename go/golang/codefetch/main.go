package main

import "fmt"

func main() {
	// golang默认fetch工具不支持500MB以上的repo，go.mod大小限制为16MB
	fmt.Println(16 << 20)
	fmt.Println(500 << 20)
	fmt.Println(500 * 1024 * 1024)
	fmt.Println(1 << 10)
	fmt.Println(1 << 20)
	fmt.Println(1 << 30)
}
