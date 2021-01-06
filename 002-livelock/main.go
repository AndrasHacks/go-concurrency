package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var cadence = sync.NewCond(&sync.Mutex{})

func takeStep() {
	cadence.L.Lock()
	cadence.Wait()
	cadence.L.Unlock()
}

func tryDir(dirName string, steps *int32, out *bytes.Buffer) bool {
	fmt.Fprintf(out, "%v", dirName)
	atomic.AddInt32(steps, 1)
	takeStep()
	if atomic.LoadInt32(steps) == 1 {
		fmt.Fprintf(out, ". Scuccess!")
		return true
	}
	fmt.Fprintf(out, "%v (%d)", dirName, *steps)
	takeStep()
	atomic.AddInt32(steps, -1)
	return false
}

var left, right int32

func tryLeft(out *bytes.Buffer) bool {
	return tryDir(" left", &left, out)
}

func tryRight(out *bytes.Buffer) bool {
	return tryDir(" right", &right, out)
}

func walk(people *sync.WaitGroup, name string) {
	var out bytes.Buffer
	defer func() { fmt.Println(out.String()) }()
	defer people.Done()
	fmt.Fprintf(&out, "%v is trying to scoot", name)
	for i := 0; i < 5; i++ {
		if tryLeft(&out) || tryRight(&out) {
			continue
		}
	}
	fmt.Fprintf(&out, "\n%v is shaking her head!\n", name)

}

func main() {
	// Set up cadence of operations
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
