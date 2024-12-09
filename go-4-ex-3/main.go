package main

// TODO: implement the function computeQuadraticFormula
import (
	"fmt"
	"math"
)

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiscriminant(a, b, c)
	var solutions []float64

	if D > 0 {
		solution1 := (-b + math.Sqrt(D)) / (2 * a)
		solution2 := (-b - math.Sqrt(D)) / (2 * a)
		solutions = append(solutions, solution1, solution2)
	} else if D == 0 {
		solution := -b / (2 * a)
		solutions = append(solutions, solution)
	}

	return solutions
}
func main() {
	// TODO: call the function computeQuadraticFormula
	testCases := []struct {
		a, b, c float64
	}{
		{3, 4, 1},
		{2, 4, 2},
		{3, 4, 2},
	}

	for _, test := range testCases {
		solutions := computeQuadraticFormula(test.a, test.b, test.c)
		fmt.Printf("Lösungen für a=%.1f, b=%.1f, c=%.1f: %v\n", test.a, test.b, test.c, solutions)
	}
}
//Nicolas