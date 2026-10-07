package main

// time GOMAXPROCS=1 go run 04_why_goroutines.go

import (
	"fmt"
	"time"
)

func check(host string, results chan string) {
	time.Sleep(time.Second)
	results <- host + " ok"
	close(results)
}

func main() {
	results := make(chan string)

	hosts := []string{"srv1", "srv2", "srv3"}

	for _, h := range hosts {
		go check(h, results)
	}

	for v := range results {
		fmt.Println(v)
	}
}
