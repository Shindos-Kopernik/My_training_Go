package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go say("hello")
}
func say(word string, ch chan int) {
	fmt.Println(word)
	ch <- 1
}
