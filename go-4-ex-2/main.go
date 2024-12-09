package main

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
import (
    "fmt"
    "math"
)

func computeHypotenuse(a, b float64) float64 {
    return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
    a float64
    b float64
}

func (s ShortSides) Hypotenuse() float64 {
    return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println("Test computeHypotenuse:")
    fmt.Println(computeHypotenuse(3, 4))
    fmt.Println(computeHypotenuse(5, 12))
    fmt.Println(computeHypotenuse(8, 15))

    fmt.Println("\nTest Hypotenuse Methode:")
    sides1 := ShortSides{3, 4}
    fmt.Println(sides1.Hypotenuse())

    sides2 := ShortSides{5, 12}
    fmt.Println(sides2.Hypotenuse())

    sides3 := ShortSides{8, 15}
    fmt.Println(sides3.Hypotenuse())
}
//Nicolas
