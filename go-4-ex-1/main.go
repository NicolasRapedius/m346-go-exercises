package main

// TODO: implement the function computeGrade
import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if gotPoints < 0 || maxPoints <= 0 || gotPoints > maxPoints {
		return 0, errors.New("ungültige Punktzahlen")
	}

	percentage := gotPoints / maxPoints
	grade := 1 + 5*percentage

	return grade, nil
}

func main() {
	// TODO: call the function computeGrade
	if grade, err := computeGrade(17.5, 28.0); err == nil {
		fmt.Println(grade)
	} else {
		fmt.Println(err)
	}

	if grade, err := computeGrade(25.0, 30.0); err == nil {
		fmt.Println(grade)
	} else {
		fmt.Println(err)
	}

	if grade, err := computeGrade(0.0, 20.0); err == nil {
		fmt.Println(grade)
	} else {
		fmt.Println(err)
	}

	if grade, err := computeGrade(-5.0, 20.0); err == nil {
		fmt.Println(grade)
	} else {
		fmt.Println(err)
	}

	if grade, err := computeGrade(25.0, 20.0); err == nil {
		fmt.Println(grade)
	} else {
		fmt.Println(err)
	}
}
//Nicolas
