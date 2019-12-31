package main

import (
	"fmt"
	"log"
)

const inputPath = "input.txt"

func main() {
	inputs, err := Read(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fuel sum: ", Sum(inputs))
}
