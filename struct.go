package main

type Person struct {
	id   int
	age  int
	name string
}

// Receiver functions
func (person Person) toString() string {
	return "Hi from toString"
}

func main() {
	person := Person{id: 10, name: "Gregori", age: 24}
	person.name = "Piñeres"
	println(person.toString())
}
