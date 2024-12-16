package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func convertCelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func convertFahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return convertCelsiusToFahrenheit(c)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return convertFahrenheitToCelsius(f)
}

func main() {
c1 := Celsius(0.0)
f1 := convertCelsiusToFahrenheit(c1)

c2 := Celsius(100.0)
f2 := convertCelsiusToFahrenheit(c2)

c3 := Celsius(23.0)
f3 := convertCelsiusToFahrenheit(c3)

fmt.Println("Rückumrechnung:")
fmt.Printf("Fahrenheit von %.1f°C: %.1f°F\n", c1, f1)
fmt.Printf("Celsius von %.1f°F: %.1f°C\n", f1, convertFahrenheitToCelsius(f1))

fmt.Printf("Fahrenheit von %.1f°C: %.1f°F\n", c2, f2)
fmt.Printf("Celsius von %.1f°F: %.1f°C\n", f2, convertFahrenheitToCelsius(f2))

fmt.Printf("Fahrenheit von %.1f°C: %.1f°F\n", c3, f3)
fmt.Printf("Celsius von %.1f°F: %.1f°C\n", f3, convertFahrenheitToCelsius(f3))

var cozy Celsius = 23.0
fmt.Printf("Cozy in Fahrenheit: %.1f°F\n", cozy.ConvertToFahrenheit())

var cold Fahrenheit = 15.3
fmt.Printf("Cold in Celsius: %.1f°C\n", cold.ConvertToCelsius())
}
//Nicolas
