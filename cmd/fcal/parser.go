package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Lateks/cotsworth/cal"
)

type command struct {
	numMonths               int
	firstMonth              *cal.IFCDate
	highlightDay            *cal.IFCDate
	showRelationToGregorian bool
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
		return cal.January, fmt.Errorf("invalid month value %d (use %d-%d)", month, 1, cal.MonthsInYear)
	}
	return cal.IFCMonth(month), nil
}

func parseGregorianMonth(arg string) (time.Month, error) {
	month, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		for i, monthName := range cal.LongMonthNames {
			if monthName == "Sol" {
				continue
			}
			if strings.ToLower(monthName) == strings.ToLower(arg) {
				if i > 7 {
					return time.Month(i - 1), nil
				} else {
					return time.Month(i), nil
				}
			}
		}
	}
	if month < 1 || month > 12 {
		return time.January, fmt.Errorf("invalid Gregorian month value: %d (use 1-12)", month)
	}
	return time.Month(int(month)), nil
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

func parseGregorianDay(arg string, month time.Month, year int) (int, error) {
	day, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		return 0, err
	}
	if day < 1 {
		return 0, fmt.Errorf("invalid Gregorian day value: %d", day)
	}

	switch month {
	case time.January:
		if cal.IsLeapYear(year) && day > 29 {
			return 0, fmt.Errorf("invalid day in month %s: %d (use 1-29)", month, day)
		}
		if day > 28 {
			return 0, fmt.Errorf("invalid day in month %s: %d (use 1-28)", month, day)
		}
	case time.April, time.June, time.September, time.November:
		if day > 30 {
			return 0, fmt.Errorf("invalid day in month %s: %d (use 1-30)", month, day)
		}
	default:
		if day > 31 {
			return 0, fmt.Errorf("invalid day in month %s: %d (use 1-31)", month, day)
		}
	}

	return int(day), nil
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
	if flags.ParseGregorian && argCount < 3 {
		log.Fatalln("Gregorian parsing mode expects year, month and day parameters")
	}

	switch argCount {
	case 3:
		var year, day int
		var err error
		if year, err = parseYear(args[0]); err != nil {
			logArgParseError(err, args[0])
		}
		if flags.ParseGregorian {
			var month time.Month
			if month, err = parseGregorianMonth(args[1]); err != nil {
				logArgParseError(err, args[1])
			}
			if day, err = parseGregorianDay(args[2], month, year); err != nil {
				logArgParseError(err, args[2])
			}
			monthSelection = cal.DateAt(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
		} else {
			var month cal.IFCMonth
			if month, err = parseMonth(args[1]); err != nil {
				logArgParseError(err, args[1])
			}
			if day, err = parseDay(args[2], month, year); err != nil {
				logArgParseError(err, args[2])
			}
			monthSelection = cal.NewIFCDate(year, month, day)
		}
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
		numMonths:               numMonthsToShow,
		firstMonth:              startMonth,
		highlightDay:            highlightDay,
		showRelationToGregorian: flags.ShowRelationToGregorian,
	}
}
