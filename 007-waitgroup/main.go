package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("Main is beginning!")

	// the same WaitGroup instance needs to be passed around with pointers.
	wg := sync.WaitGroup{}
	num := 1e4
	wg.Add(int(num))
	for i := 0; float64(i) < num; i++ {
		go async(&wg)
		fmt.Println(i)
	}
	wg.Wait()
	fmt.Println("Main is done!")

}

func async(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Imitating some work to be done!")
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	fmt.Printf("My goroutine id: %d\n", id)
	fmt.Println(string(buf[:n]))
	time.Sleep(time.Second * 1)
}
