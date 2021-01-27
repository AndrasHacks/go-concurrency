package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	var m sync.RWMutex
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := math.Pow(2, float64(i))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			int(count),
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}

func test(count float64, mutex, rwMutex sync.Locker) time.Duration {
	wg := sync.WaitGroup{}
	wg.Add(int(count) + 1)
	begin := time.Now()
	go producer(&wg, &mutex)
	for i := 0; i < int(count); i++ {
		go consumer(&wg, &rwMutex)
	}
	wg.Wait()
	return time.Since(begin)
}

func producer(wg *sync.WaitGroup, locker *sync.Locker) {
	// fmt.Println("In producer...")
	// defer func() { fmt.Println("Exit producer...") }()
	defer wg.Done()
	(*locker).Lock()
	defer (*locker).Unlock()
	time.Sleep(1)
}

func consumer(wg *sync.WaitGroup, locker *sync.Locker) {
	// fmt.Println("In consumer...")
	// defer func() { fmt.Println("Exit consumer...") }()
	defer wg.Done()
	(*locker).Lock()
	defer (*locker).Unlock()
}
