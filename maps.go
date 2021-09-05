package main

import "fmt"

type Friend struct {
	name string
	age int
}

func main() {
    m := make(map[string]string)
    m["name"] = "Gregori"
    m["address"] = "Mérida, Venezuela"

    fmt.Println("Hi", m["name"] + "!")
    fmt.Println("Are you from", m["address"] + "?")

    f := []Friend {
    	{name: "Maria", age: 79},
    	{name: "Javier", age: 54},
    }

    eachPrint(f)
}

func eachPrint(friends []Friend) {
	for _, friend := range friends {
		println(friend.name)
		println(friend.age)
	}
}
