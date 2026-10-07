package main

import "fmt"

// This introduces buffer
func main() {
	ch := make(chan int, 2)
	// ch := make(chan int)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
