package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Println("Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	diceLogFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Println("Error creating dice.log:", err)
		return
	}
	defer diceLogFile.Close()

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(diceLogFile, "the dice was rolled at, when")
	// go run ex3/main.go TODO
	// Nicolas
}
