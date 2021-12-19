package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Lateks/cotsworth/cal"
	fcalFmt "github.com/Lateks/cotsworth/fmt"
)

type Flags struct {
	ParseGregorian        bool
	ShowSurroundingMonths int
}

func displayMonth(monthDate *cal.IFCDate, highlightDate *cal.IFCDate) {
	monthLines := fcalFmt.MonthToLines(monthDate.Year, monthDate.Month, highlightDate)
	fmt.Println(strings.Join(monthLines, "\n"))
}

func displayMonthsOnLine(numMonths int, startMonth *cal.IFCDate, highlightDate *cal.IFCDate) {
	if numMonths < 1 {
		return
	}
	if numMonths == 1 {
		displayMonth(startMonth, highlightDate)
		return
	}

	months := make([][]string, numMonths)
	month := startMonth
	for m := 0; m < numMonths; m++ {
		months[m] = fcalFmt.MonthToLines(month.Year, month.Month, highlightDate)
		month = month.PlusMonths(1)
	}

	for i := 0; i < len(months[0]); i++ {
		for j := 0; j < numMonths-1; j++ {
			fmt.Print(months[j][i])
		}
		fmt.Println(months[numMonths-1][i])
	}
}

const maxMonthsPerLine = 3

func displayMonths(numMonths int, startMonth *cal.IFCDate, highlightDate *cal.IFCDate) {
	for numMonths > 0 {
		monthsToDisplay := int(math.Min(maxMonthsPerLine, float64(numMonths)))
		displayMonthsOnLine(monthsToDisplay, startMonth, highlightDate)
		startMonth = startMonth.PlusMonths(monthsToDisplay)
		numMonths -= monthsToDisplay
	}
}

func Execute(flags *Flags, args []string) {
	command := parseArgs(flags, args)
	displayMonths(command.numMonths, command.firstMonth, command.highlightDay)
}
