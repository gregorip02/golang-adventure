package main

func main() {
	inicial := 10
	println(inicial);

	final := &inicial
	println(final)  // Imprime la dirección en memoria de la variable inicial
	println(*final) // Imprime el valor de la variable inicial
}
