package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if err := exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exec() error {
	go schedule("Budapest")
	go schedule("London")
	for true {
		fmt.Printf("\nSome comutation happening in the main goroutine...\n")
		time.Sleep(2 * time.Second)
	}
	return nil
}

func schedule(city string) {
	for true {
		currentTime := make(chan string, 1)
		url := fmt.Sprintf("http://worldtimeapi.org/api/timezone/Europe/%v.txt", city)
		go fetchTime(currentTime, url)
		fmt.Printf("Current time in %v: %q\n", city, <-currentTime)
		time.Sleep(3 * time.Second)
	}
}

func fetchTime(ch chan string, url string) {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	sb := string(body)
	lines := strings.Split(sb, "\n")
	ch <- lines[2]
}
