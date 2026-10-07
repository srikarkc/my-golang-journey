package main

import "fmt"

func return_one(i int, ch chan int) {
	ch <- i
}

func main() {
	ch := make(chan int)

	go return_one(1, ch)

	fmt.Println(<-ch)
}
