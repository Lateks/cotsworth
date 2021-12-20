package cal_test

import (
	"testing"
	"time"

	"github.com/Lateks/cotsworth/cal"
)

func TestIFCDateConversion(t *testing.T) {
	tzHelsinki, err := time.LoadLocation("Europe/Helsinki")
	if err != nil {
		t.Fatalf("Error loading Helsinki timezone")
	}

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
		{
			time.Date(2021, time.January, 1, 1, 00, 0, 0, tzHelsinki),
			cal.NewIFCDate(2021, cal.January, 1),
		},
	} {
		date := cal.DateAt(input.time)
		if !date.Equal(input.ifcDate) {
			t.Errorf("%d: Expected %+v but found %+v\n", i, input.ifcDate, date)
		}

		timeStamp := date.ToUTCTime()
		y, m, d := input.time.Date()
		dateWithoutTime := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		if !timeStamp.Equal(dateWithoutTime) {
			t.Errorf("%d: Expected %+v but found %+v\n", i, dateWithoutTime, timeStamp)
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

func TestAddMonthsToDate(t *testing.T) {
	for i, input := range []struct {
		ifcDate     *cal.IFCDate
		monthsToAdd int
		result      *cal.IFCDate
	}{
		{
			cal.NewIFCDate(2020, cal.January, 1),
			13,
			cal.NewIFCDate(2021, cal.January, 1),
		},
		{
			cal.NewIFCDate(2020, cal.Sol, 1),
			13,
			cal.NewIFCDate(2021, cal.Sol, 1),
		},
		{
			cal.NewIFCDate(2021, cal.December, 1),
			1,
			cal.NewIFCDate(2022, cal.January, 1),
		},
		{
			cal.NewIFCDate(2021, cal.December, 29),
			1,
			cal.NewIFCDate(2022, cal.January, 28),
		},
		{
			cal.NewIFCDate(2021, cal.December, 29),
			0,
			cal.NewIFCDate(2021, cal.December, 29),
		},
		{
			cal.NewIFCDate(2020, cal.June, 29),
			1,
			cal.NewIFCDate(2020, cal.Sol, 28),
		},
		{
			cal.NewIFCDate(2020, cal.December, 1),
			0,
			cal.NewIFCDate(2020, cal.December, 1),
		},
	} {
		newDate := input.ifcDate.PlusMonths(input.monthsToAdd)
		if !newDate.Equal(input.result) {
			t.Errorf("%d: Expected %+v but found %+v\n", i, input.result, newDate)
		}
	}
}
func TestMinusMonthsFromDate(t *testing.T) {
	for i, input := range []struct {
		ifcDate          *cal.IFCDate
		monthsToSubtract int
		result           *cal.IFCDate
	}{
		{
			cal.NewIFCDate(2020, cal.January, 1),
			1,
			cal.NewIFCDate(2019, cal.December, 1),
		},
		{
			cal.NewIFCDate(2020, cal.Sol, 1),
			13,
			cal.NewIFCDate(2019, cal.Sol, 1),
		},
		{
			cal.NewIFCDate(2021, cal.January, 28),
			1,
			cal.NewIFCDate(2020, cal.December, 28),
		},
		{
			cal.NewIFCDate(2020, cal.June, 29),
			1,
			cal.NewIFCDate(2020, cal.May, 28),
		},
		{
			cal.NewIFCDate(2020, cal.December, 29),
			0,
			cal.NewIFCDate(2020, cal.December, 29),
		},
	} {
		newDate := input.ifcDate.MinusMonths(input.monthsToSubtract)
		if !newDate.Equal(input.result) {
			t.Errorf("%d: Expected %+v but found %+v\n", i, input.result, newDate)
		}
	}
}
