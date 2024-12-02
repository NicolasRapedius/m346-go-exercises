package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	for i := 2; i < 5; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	// TODO: correct up to index 4 using direct element access

	fibs = append(fibs, fibs[3]+fibs[4]) // TODO: replace 0 with the next Fibonacci number
	// TODO: compute three more Fibonacci numbers and append them
	for i := 0; i < 3; i++ {
		nextFib := fibs[len(fibs)-1] + fibs[len(fibs)-2]
		fibs = append(fibs, nextFib)
	}
	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
	//Nicolas
}
