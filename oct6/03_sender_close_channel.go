package main

import "fmt"

func sendNums(count int, ch chan int) {
	for i := 1; i <= count; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go sendNums(5, ch)

	for v := range ch {
		fmt.Println(v)
	}
}
