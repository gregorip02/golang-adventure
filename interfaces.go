package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Jsonable interface {
	toJson() ([]byte, error)
}

func printJson(contract Jsonable) {
	data, err := contract.toJson()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println(string(data))
	}
}

type Car struct {
	Year int `json:"year"`
	Model string `json:"model"`
}

func (car Car) toJson() ([]byte, error) {
	return json.Marshal(car)
}

func main() {
	var car Car = Car{Year: 2021, Model: "Tesla S"}
	printJson(car)
}
