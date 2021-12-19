package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Lateks/cotsworth/cal"
)

type command struct {
	numMonths    int
	firstMonth   *cal.IFCDate
	highlightDay *cal.IFCDate
}

func parseYear(arg string) (int, error) {
	year, err := strconv.ParseInt(arg, 10, 32)
	if year < 0 {
		return 0, fmt.Errorf("invalid year value: %d", year)
	}
	return int(year), err
}

func parseMonth(arg string) (cal.IFCMonth, error) {
	month, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		for i, monthName := range cal.LongMonthNames {
			if strings.ToLower(monthName) == strings.ToLower(arg) {
				return cal.IFCMonth(i + 1), nil
			}
		}
		return cal.January, err
	}
	if month < 1 || month > cal.MonthsInYear {
		return cal.January, fmt.Errorf("invalid month value: %d", month)
	}
	return cal.IFCMonth(month), nil
}

func parseDay(arg string, month cal.IFCMonth, year int) (int, error) {
	day, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		return 0, err
	}
	if day < 1 || ((month == cal.December) || (month == cal.June && cal.IsLeapYear(year))) && day > 29 || day > 28 {
		return 0, fmt.Errorf("invalid day value: %d", day)
	}

	return int(day), err
}

func logArgParseError(err error, arg string) {
	log.Fatalf("Error parsing argument %s: %s\n", arg, err)
}

func parseArgs(flags *Flags, args []string) *command {
	today := cal.DateAt(time.Now())
	monthSelection := today
	highlightDay := today
	numMonthsToShow := 1 + flags.ShowSurroundingMonths

	argCount := len(args)
	switch argCount {
	case 3:
		var year, day int
		var month cal.IFCMonth
		var err error
		if year, err = parseYear(args[0]); err != nil {
			logArgParseError(err, args[0])
		}
		if month, err = parseMonth(args[1]); err != nil {
			logArgParseError(err, args[1])
		}
		if day, err = parseDay(args[2], month, year); err != nil {
			logArgParseError(err, args[2])
		}
		monthSelection = cal.NewIFCDate(year, month, day)
		highlightDay = monthSelection
	case 2:
		var year int
		var month cal.IFCMonth
		var err error
		if year, err = parseYear(args[0]); err != nil {
			logArgParseError(err, args[0])
		}
		if month, err = parseMonth(args[1]); err != nil {
			logArgParseError(err, args[1])
		}
		monthSelection = cal.NewIFCDate(year, month, 1)
	case 1:
		// Try to parse argument as a year.
		year, err := parseYear(args[0])
		if err == nil {
			monthSelection = cal.NewIFCDate(year, cal.January, 1)
			numMonthsToShow = cal.MonthsInYear + flags.ShowSurroundingMonths
		} else {
			// Failing that, assume it's a month.
			month, err := parseMonth(args[0])
			if err != nil {
				logArgParseError(err, args[0])
			}
			monthSelection = cal.NewIFCDate(today.Year, month, 1)
		}
	}

	startMonth := monthSelection.MinusMonths(flags.ShowSurroundingMonths / 2)
	return &command{
		numMonths:    numMonthsToShow,
		firstMonth:   startMonth,
		highlightDay: highlightDay,
	}
}
