package main

import "fmt"

func say(msg string, done chan bool) {
	fmt.Println(msg)
	done <- true // send: I'm finished
}

func main() {
	done := make(chan bool)

	go say("hello", done)
	<-done // receive: main waits here until a value arrives
	fmt.Println("world")
}
