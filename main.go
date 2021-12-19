package main

import (
	"flag"
	"github.com/Lateks/cotsworth/cmd"
)

func main() {
	var monthsToDisplay int
	flag.IntVar(&monthsToDisplay, "n", 1, "number of months to display")
	flag.Parse()

	flags := &cmd.Flags{
		ShowSurroundingMonths: monthsToDisplay - 1,
	}

	cmd.Execute(flags, flag.Args())
}
