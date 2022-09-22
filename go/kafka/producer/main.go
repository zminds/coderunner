package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("Hello World!")
	cfg := sarama.NewConfig()
	fmt.Println(cfg)

}
