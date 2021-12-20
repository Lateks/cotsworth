package cal

import (
	"time"
)

const (
	DaysInWeek   = 7
	daysInMonth  = 28
	WeeksInMonth = 4
	DaysInYear   = 365
	MonthsInYear = 13
	LeapDayDate  = 6*daysInMonth + 1 // Leap day occurs on June 29.
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

var LongMonthNames = []string{
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
	return LongMonthNames[int(m)-1]
}

func toIFCMonth(zeroBasedVal int) IFCMonth {
	newVal := (zeroBasedVal % MonthsInYear) + 1
	if newVal < 1 {
		newVal = MonthsInYear + newVal
	}
	return IFCMonth(newVal)
}

func (m IFCMonth) PlusMonths(months int) IFCMonth {
	return toIFCMonth(int(m) - 1 + months)
}

type Weekday int

const (
	Sunday Weekday = iota
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

var shortWeekdayNames = []string{
	"Su",
	"Mo",
	"Tu",
	"We",
	"Th",
	"Fr",
	"Sa",
	"LD",
	"YD",
}

func (wd Weekday) String() string {
	return longWeekdayNames[wd]
}

func (wd Weekday) ShortFormat() string {
	return shortWeekdayNames[wd]
}

type IFCDate struct {
	Year      int
	Month     IFCMonth
	Day       int
	dayOfYear int
}

func NewIFCDate(year int, month IFCMonth, day int) *IFCDate {
	dayOfYear := (int(month)-1)*daysInMonth + day
	if IsLeapYear(year) && month > June {
		dayOfYear++
	}

	return &IFCDate{
		Year:      year,
		Month:     month,
		Day:       day,
		dayOfYear: dayOfYear,
	}
}

func (d *IFCDate) ToUTCTime() time.Time {
	date := time.Date(d.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
	return date.Add(time.Duration(d.dayOfYear-1) * 24 * time.Hour)
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

	weekday := (d.Day - 1) % DaysInWeek
	return Weekday(weekday)
}

func (d *IFCDate) PlusMonths(months int) *IFCDate {
	if months == 0 {
		return d
	}

	var day int
	if d.Day == 29 {
		day = 28
	} else {
		day = d.Day
	}
	month := d.Month.PlusMonths(months)

	changeInYears := months / MonthsInYear
	remMonths := months % MonthsInYear
	if int(d.Month)+remMonths > MonthsInYear {
		changeInYears++
	} else if int(d.Month)+remMonths < 1 {
		changeInYears--
	}
	year := d.Year + changeInYears

	return NewIFCDate(year, month, day)
}

func (d *IFCDate) MinusMonths(months int) *IFCDate {
	return d.PlusMonths(-months)
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

func IsLeapYear(year int) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	}

	return year%4 == 0
}

func date(t time.Time) (year int, month IFCMonth, day int, dayOfYear int) {
	year = t.Year()
	dayOfGregorianYear := t.YearDay()
	dayOfYear = dayOfGregorianYear

	if IsLeapYear(year) {
		if dayOfGregorianYear == LeapDayDate {
			month = June
			day = 29
			return
		}

		// Ignore leap day after it has gone by.
		if dayOfGregorianYear > LeapDayDate {
			dayOfGregorianYear--
		}
	}

	if dayOfGregorianYear == DaysInYear {
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

func DaysInMonth(year int, month IFCMonth) int {
	switch month {
	case June:
		if IsLeapYear(year) {
			return 29
		}
	case December:
		return 29
	}
	return 28
}
