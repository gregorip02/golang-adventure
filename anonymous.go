package main

import (
	"fmt"
)

func main() {
	x := 100

	sum := func (a, b int) int {
		return a + b
	}

	func() {
		// Esta función tiene un alcance fuera de su scope.
		// Es decir, puede modificar el valor de x.
		x = x * 2
	}()

	fmt.Println(x) // Imprime 200

	fmt.Println(sum(10, 20))
}
