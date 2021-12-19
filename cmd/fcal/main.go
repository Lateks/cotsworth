package main

import (
	"flag"
)

func main() {
	var gregorian bool
	var monthsToDisplay int
	flag.IntVar(&monthsToDisplay, "n", 1, "number of months to display")
	flag.BoolVar(&gregorian, "g", false, "parse the parameters as Gregorian calendar date (requires year, month and day parameters)")
	flag.Parse()

	flags := &Flags{
		ParseGregorian:        gregorian,
		ShowSurroundingMonths: monthsToDisplay - 1,
	}

	Execute(flags, flag.Args())
}
