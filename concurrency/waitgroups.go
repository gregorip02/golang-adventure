package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// wg.Add incrementa en 1 la cantidad de rutinas agregadas al wg.
		wg.Add(1)
		go sleeping(i, &wg)
	}

	wg.Wait()
}

func sleeping(index int, wg *sync.WaitGroup) {
	// defer se ejecuta cuando la rutina es terminada.
	// wg.Done derementa en 1 la cantidad de rutinas agregadas al wg.
	defer wg.Done()
	fmt.Printf("[sleeping]: starting %d", index)
	time.Sleep(2 * time.Second)
	fmt.Printf("[sleeping]: done %d", index)
}
