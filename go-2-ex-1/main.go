package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			FirstName: "Nicolas",
			LastName:  "Rapedius",
		},
		BirthDate: BirthDate{
			Day:   24,
			Month: 4,
			Year:  2008,
		},
		NumberOfSiblings: 1,   // TODO: adjust
		ZodiacSign:       '\u2649', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
