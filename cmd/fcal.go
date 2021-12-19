package cmd

import (
	"fmt"
	"github.com/Lateks/cotsworth/cal"
	fcalFmt "github.com/Lateks/cotsworth/fmt"
	"log"
	"strings"
)

func displayMonth(monthDate *cal.IFCDate, highlightDate *cal.IFCDate) {
	monthLines := fcalFmt.MonthToLines(monthDate.Year, monthDate.Month, highlightDate)
	fmt.Println(strings.Join(monthLines, "\n"))
}

func Execute(args []string) {
	command := parseArgs(args)
	if command.numMonths == 1 {
		displayMonth(command.firstMonth, command.highlightDay)
	} else {
		log.Fatalf("Not implemented")
	}
}
