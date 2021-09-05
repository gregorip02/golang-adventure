package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declaración de una variable explicitamente.
	var x int
	x = 8

	// Declaración y asignación de una variable implicitamente.
	// El tipo de dato es inferido a partir del dato asignado.
	y := 7

	fmt.Println(x) // Imprime 8
	fmt.Println(y) // Imprime 7

	// En go no existen las excepciones, en su lugar, debemos capturar
	// errores explicitamente.
	const str string = "10"
	data, err := strconv.ParseInt(str, 0, 64)

	// Notas: ¿Como puedo implementar un operador ternario en go?
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	// Maps: Los maps son estructuras de datos clave -> valor.
	m := make(map[string]int)
	m["age"] = 24
	fmt.Println(m["age"])

	// Slices: Los slices son similares a listas iterables de datos.
	s := []int{1, 2, 3}
	each(s)
	s = append(s, 16)
	each(s)
}

func each(s []int) {
	for index, value := range s {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
