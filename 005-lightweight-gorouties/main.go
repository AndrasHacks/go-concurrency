package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main() {
	if err := exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exec() error {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()

	for i := numGoroutines; i > 1; i-- {
		go noop()
	}
	wg.Done()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
	return nil
}
