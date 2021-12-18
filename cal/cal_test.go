package cal_test

import (
	"testing"
	"time"

	"github.com/Lateks/cotsworth/cal"
)

func TestUTCDates(t *testing.T) {
	for i, input := range []struct {
		time    time.Time
		ifcDate *cal.IFCDate
	}{
		{
			time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			cal.NewIFCDate(1970, cal.January, 1),
		},
		{
			time.Date(2021, time.December, 18, 19, 40, 0, 0, time.UTC),
			cal.NewIFCDate(2021, cal.December, 16),
		},
		{
			time.Date(2020, time.June, 17, 12, 00, 0, 0, time.UTC),
			cal.NewIFCDate(2020, cal.June, 29),
		},
		{
			time.Date(2020, time.June, 18, 12, 00, 0, 0, time.UTC),
			cal.NewIFCDate(2020, cal.Sol, 1),
		},
		{
			time.Date(2021, time.December, 31, 12, 00, 0, 0, time.UTC),
			cal.NewIFCDate(2021, cal.December, 29),
		},
	} {
		date := cal.DateAt(input.time)
		if !date.Equal(input.ifcDate) {
			t.Errorf("%d: Expected %+v but found %+v\n", i, input.ifcDate, date)
		}
	}
}

func TestWeekdays(t *testing.T) {
	for i, input := range []struct {
		ifcDate *cal.IFCDate
		weekday cal.Weekday
	}{
		{
			cal.NewIFCDate(2000, cal.January, 1),
			cal.Sunday,
		},
		{
			cal.NewIFCDate(1990, cal.August, 13),
			cal.Friday,
		},
		{
			cal.NewIFCDate(2021, cal.December, 28),
			cal.Saturday,
		},
		{
			cal.NewIFCDate(2021, cal.December, 29),
			cal.YearDay,
		},
		{
			cal.NewIFCDate(2022, cal.January, 1),
			cal.Sunday,
		},
		{
			cal.NewIFCDate(2020, cal.June, 28),
			cal.Saturday,
		},
		{
			cal.NewIFCDate(2020, cal.June, 29),
			cal.LeapDay,
		},
		{
			cal.NewIFCDate(2020, cal.Sol, 1),
			cal.Sunday,
		},
	} {
		weekday := input.ifcDate.Weekday()
		if weekday != input.weekday {
			t.Errorf("%d: Expected %s but found %s\n", i, input.weekday, weekday)
		}
	}
}
