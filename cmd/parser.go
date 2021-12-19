package cmd

import (
	"fmt"
	"github.com/Lateks/cotsworth/cal"
	"log"
	"strconv"
	"strings"
	"time"
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

func parseArgs(args []string) *command {
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
		date := cal.NewIFCDate(year, month, day)
		command := &command{
			numMonths:    1,
			firstMonth:   date,
			highlightDay: date,
		}
		return command
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
		date := cal.NewIFCDate(year, month, 1)
		command := &command{
			numMonths:    1,
			firstMonth:   date,
			highlightDay: nil,
		}
		return command
	case 1:
		month, err := parseMonth(args[0])
		if err != nil {
			logArgParseError(err, args[0])
		}
		year := time.Now().Year()
		date := cal.NewIFCDate(year, month, 1)
		command := &command{
			numMonths:    1,
			firstMonth:   date,
			highlightDay: nil,
		}
		return command
	}

	today := cal.DateAt(time.Now())
	return &command{
		numMonths:    1,
		firstMonth:   today,
		highlightDay: today,
	}
}
