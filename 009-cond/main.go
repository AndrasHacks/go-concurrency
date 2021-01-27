package main

import (
	"fmt"
	"sync"
)

var queue = make([]interface{}, 0, 10)

func main() {
	cond := sync.Cond{L: &sync.Mutex{}}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go producer(&cond, &wg)
	go consumer(&cond, &wg)
	wg.Wait()

}

func producer(cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for len(queue) == 2 {
			cond.Wait()
		}
		queue = append(queue, struct{}{})
		fmt.Printf("Queued %d. item\n", i+1)
		cond.L.Unlock()
		cond.Signal()
	}
}

func consumer(cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for len(queue) == 0 {
			cond.Wait()
		}
		// simulate dequeue
		queue = queue[1:]
		fmt.Printf("Dequeued %d. item\n", i+1)
		cond.L.Unlock()
		cond.Signal()
	}
}
