package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	go generator(in)
	go double(in, out)
	print(out)
}

func print(channel chan int) {
	for value := range channel {
		fmt.Println(value)
	}
}

func double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
		// in es un canal de lectura, no podemos enviar data en el.
		// in <- 1
	}
	close(out)
}

func generator(channel chan<- int) {
	for i := 0; i <= 10; i++ {
		channel <- i
	}
	close(channel)
}
