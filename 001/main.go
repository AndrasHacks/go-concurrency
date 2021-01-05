package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// unsafe()
	// notIdiomatic()
	deadlock()
}

func deadlock() {
	type deadlocker struct {
		data   int
		locker sync.Mutex
	}

	var wg sync.WaitGroup
	sum := func(v1, v2 *deadlocker) int {
		defer wg.Done()
		v1.locker.Lock()
		p1 := v1.data
		defer v1.locker.Unlock()

		// Some additional computation is happening here!
		time.Sleep(5 * time.Second)

		v2.locker.Lock()
		p2 := v2.data
		defer v2.locker.Unlock()

		return p1 + p2
	}

	wg.Add(2)
	a := deadlocker{data: 3}
	b := deadlocker{data: 2}
	go sum(&a, &b)
	go sum(&b, &a)
	wg.Wait()
}

func unsafe() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("Data equals %d!\n", data)
		time.Sleep(1 * time.Second)
		fmt.Printf("Data equals %d!\n", data)
	}
}

func notIdiomatic() {
	var lock sync.Mutex
	var data int
	go func() {
		lock.Lock()
		data++
		lock.Unlock()
	}()
	time.Sleep(1 * time.Second)
	lock.Lock()
	if data == 0 {
		fmt.Println("Data equals zero!")
	} else {
		fmt.Printf("Data equals %d!\n", data)
	}
	lock.Unlock()
}
