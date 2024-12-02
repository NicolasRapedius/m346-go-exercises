package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	type Module struct {
		Number int
		Name   string
	}
	// TODO: output everything using fmt.Println()
	type ClassModules map[Module][]Class

	studentsClassA := []Student{
		{"Nicolas", "Rapedius"},
		{"Anna", "Müller"},
		{"Tom", "Schmidt"},
	}

	studentsClassB := []Student{
		{"Lisa", "Meier"},
		{"Paul", "Schneider"},
		{"Sophie", "Fischer"},
	}

	classA := Class{Name: "Klasse A", Students: studentsClassA}
	classB := Class{Name: "Klasse B", Students: studentsClassB}

	module1 := Module{Number: 101, Name: "Mathematik"}
	module2 := Module{Number: 102, Name: "Informatik"}
	module3 := Module{Number: 103, Name: "Physik"}

	classModules := ClassModules{
		module1: {classA, classB},
		module2: {classA},
		module3: {classB},
	}

	fmt.Println("Klassen und ihre Schüler:")
	for _, class := range []Class{classA, classB} {
		fmt.Printf("Name: %s\n", class.Name)
		for _, student := range class.Students {
			fmt.Printf("  - %s %s\n", student.FirstName, student.LastName)
		}
	}

	fmt.Println("\nModule und die besuchenden Klassen:")
	for module, classes := range classModules {
		fmt.Printf("Modul %d: %s\n", module.Number, module.Name)
		for _, class := range classes {
			fmt.Printf("  - %s\n", class.Name)
		}
	}
}
