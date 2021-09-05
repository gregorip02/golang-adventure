package main

import (
	"fmt"
	"time"
)

func main() {
	// La función make(t type) construye un objeto de tipo map, slice o chan.
	// el canal creado nos ayuda a que el proceso principal del programa: main,
	// espere hasta el final de la ejecución de la rutina sleeping.
	channel := make(chan bool)
	go sleeping(3, channel)
	<-channel
}

func sleeping(n time.Duration, channel chan bool) {
	time.Sleep(n * time.Second)
	fmt.Println("¡Listo!")
	channel <- true
}
