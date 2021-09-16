package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Creamos un canal con 5 espacios en el buffer.
	channel := make(chan int, 5)

	for i := 0; i < 10; i++ {
		// Agregamos una nueva sub-rutina al canal, cuando este lleno (5), se
		// bloquea el proceso hasta que tengamos mas espacio disponible para
		// agregar mas.
		channel <- 1
		wg.Add(1)
		go sleeping(i, &wg, channel)
	}

	wg.Wait()
}

func sleeping(index int, wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	fmt.Printf("[sleeping] starting %d\n", index)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Printf("[sleeping] done %d\n", index)
	<- channel
}
