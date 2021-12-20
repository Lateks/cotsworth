package main

import (
	"flag"
)

func main() {
	var relationToGregorian bool
	var gregorian bool
	var monthsToDisplay int
	flag.IntVar(&monthsToDisplay, "n", 1, "number of months to display")
	flag.BoolVar(&gregorian, "g", false, "parse the parameters as Gregorian calendar date (requires year, month and day parameters)")
	flag.BoolVar(&relationToGregorian, "r", false, "show the fixed calendar in relation to the Gregorian calendar")
	flag.Parse()

	flags := &Flags{
		ParseGregorian:          gregorian,
		ShowSurroundingMonths:   monthsToDisplay - 1,
		ShowRelationToGregorian: relationToGregorian,
	}

	Execute(flags, flag.Args())
}
