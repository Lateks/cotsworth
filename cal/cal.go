package cal

import (
	"time"
)

const (
	daysInWeek  = 7
	daysInMonth = 28
	daysInYear  = 365
	leapDay     = 6*daysInMonth + 1 // Leap day occurs on June 29.
)

type IFCMonth int

const (
	January IFCMonth = 1 + iota
	February
	March
	April
	May
	June
	Sol
	July
	August
	September
	October
	November
	December
)

var longMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"Sol",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func (m IFCMonth) String() string {
	return longMonthNames[int(m)-1]
}

type Weekday int

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday

	// LeapDay and YearDay are not part of any week.
	// They land between a Saturday and a Sunday.
	LeapDay
	YearDay
)

var longWeekdayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Leap Day",
	"Year Day",
}

func (wd Weekday) String() string {
	return longWeekdayNames[wd]
}

type IFCDate struct {
	Year      int
	Month     IFCMonth
	Day       int
	dayOfYear int
}

func NewIFCDate(year int, month IFCMonth, day int) *IFCDate {
	dayOfYear := (int(month)-1)*daysInMonth + day
	if isLeapYear(year) && month > June {
		dayOfYear++
	}

	return &IFCDate{
		Year:      year,
		Month:     month,
		Day:       day,
		dayOfYear: dayOfYear,
	}
}

func (d *IFCDate) Equal(other *IFCDate) bool {
	return d.Day == other.Day && d.Month == other.Month && d.Year == other.Year && d.dayOfYear == other.dayOfYear
}

func (d *IFCDate) IsLeapDay() bool {
	return d.Month == June && d.Day == 29
}

func (d *IFCDate) IsYearDay() bool {
	return d.Month == December && d.Day == 29
}

func (d *IFCDate) Weekday() Weekday {
	if d.IsLeapDay() {
		return LeapDay
	}
	if d.IsYearDay() {
		return YearDay
	}

	weekday := (d.Day - 1) % daysInWeek
	return Weekday(weekday)
}

func DateAt(t time.Time) *IFCDate {
	year, month, day, dayOfYear := date(t)
	return &IFCDate{
		Year:      year,
		Month:     month,
		Day:       day,
		dayOfYear: dayOfYear,
	}
}

func isLeapYear(year int) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	}

	return year%4 == 0
}

func date(t time.Time) (year int, month IFCMonth, day int, dayOfYear int) {
	year = t.Year()
	dayOfGregorianYear := t.YearDay()
	dayOfYear = dayOfGregorianYear

	if isLeapYear(year) {
		if dayOfGregorianYear == leapDay {
			month = June
			day = 29
			return
		}

		// Ignore leap day after it has gone by.
		if dayOfGregorianYear > leapDay {
			dayOfGregorianYear--
		}
	}

	if dayOfGregorianYear == daysInYear {
		// Handle year day
		month = December
		day = 29
	} else {
		monthOrdinal := dayOfGregorianYear / daysInMonth
		month = IFCMonth(monthOrdinal + 1)
		day = dayOfGregorianYear - monthOrdinal*daysInMonth
	}
	return
}
