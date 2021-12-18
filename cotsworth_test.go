package cotsworth_test

import (
	"testing"
	"time"

	"github.com/Lateks/cotsworth"
)

func TestUTCDates(t *testing.T) {
	for i, input := range []struct {
		time    time.Time
		ifcDate *cotsworth.IFCDate
	}{
		{
			time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			cotsworth.NewIFCDate(1970, cotsworth.January, 1),
		},
		{
			time.Date(2021, time.December, 18, 19, 40, 0, 0, time.UTC),
			cotsworth.NewIFCDate(2021, cotsworth.December, 16),
		},
		{
			time.Date(2020, time.June, 17, 12, 00, 0, 0, time.UTC),
			cotsworth.NewIFCDate(2020, cotsworth.June, 29),
		},
		{
			time.Date(2020, time.June, 18, 12, 00, 0, 0, time.UTC),
			cotsworth.NewIFCDate(2020, cotsworth.Sol, 1),
		},
		{
			time.Date(2021, time.December, 31, 12, 00, 0, 0, time.UTC),
			cotsworth.NewIFCDate(2021, cotsworth.December, 29),
		},
	} {
		date := cotsworth.DateAt(input.time)
		if !date.Equal(input.ifcDate) {
			t.Errorf("%d: Expected %+v but found %+v\n", i, input.ifcDate, date)
		}
	}
}

func TestWeekdays(t *testing.T) {
	for i, input := range []struct {
		ifcDate *cotsworth.IFCDate
		weekday cotsworth.Weekday
	}{
		{
			cotsworth.NewIFCDate(2000, cotsworth.January, 1),
			cotsworth.Sunday,
		},
		{
			cotsworth.NewIFCDate(1990, cotsworth.August, 13),
			cotsworth.Friday,
		},
		{
			cotsworth.NewIFCDate(2021, cotsworth.December, 28),
			cotsworth.Saturday,
		},
		{
			cotsworth.NewIFCDate(2021, cotsworth.December, 29),
			cotsworth.YearDay,
		},
		{
			cotsworth.NewIFCDate(2022, cotsworth.January, 1),
			cotsworth.Sunday,
		},
		{
			cotsworth.NewIFCDate(2020, cotsworth.June, 28),
			cotsworth.Saturday,
		},
		{
			cotsworth.NewIFCDate(2020, cotsworth.June, 29),
			cotsworth.LeapDay,
		},
		{
			cotsworth.NewIFCDate(2020, cotsworth.Sol, 1),
			cotsworth.Sunday,
		},
	} {
		weekday := input.ifcDate.Weekday()
		if weekday != input.weekday {
			t.Errorf("%d: Expected %s but found %s\n", i, input.weekday, weekday)
		}
	}
}
