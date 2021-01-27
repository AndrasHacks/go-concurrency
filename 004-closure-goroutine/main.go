package main

import (
	"fmt"
	"os"
)

func main() {
	if err := exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exec() error {
	ch := make(chan int)
	greeting := "original value"
	go func(ch chan int) {
		greeting = "modified value"
		ch <- 1
	}(ch)
	fmt.Println(<-ch)
	fmt.Println(greeting)
	return nil
}
