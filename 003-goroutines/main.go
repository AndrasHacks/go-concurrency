package main

import "fmt"

func main() {
	fmt.Println("Parent goroutine")
	ch := make(chan int)
	go concurrent(ch)
	outval := <-ch
	fmt.Println(fmt.Sprintf("Joined child go routine's result: %d", outval))
	fmt.Println("End of the parent goroutine")
}

func concurrent(ch chan int) {
	fmt.Println("Forked Child go routine")
	ch <- 1
}
