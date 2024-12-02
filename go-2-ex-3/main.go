package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Einführung in die Programmierung",
		117: "Datenbanken verwalten",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])
	// TODO: delete one
	delete(modules, 117)
	// TODO: add one
	modules[205] = "Webentwicklung mit Go"
	// TODO: replace one
	modules[346] = "Fortgeschrittene Cloud-Lösungen"
	
	fmt.Println(modules)
}
//Nicolas
