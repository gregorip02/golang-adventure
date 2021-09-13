package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool, 1)

	go async(channel)

	fmt.Println("Main is done");

	<- channel
}

func async(channel chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("Async is done");
	channel <- true
}
