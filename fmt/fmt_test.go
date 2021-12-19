package fmt_test

import (
	"github.com/Lateks/cotsworth/cal"
	"github.com/Lateks/cotsworth/fmt"
	"math"
	"testing"
)

func TestCentering(t *testing.T) {
	for i, input := range []struct {
		text       string
		fieldWidth int
		result     string
	}{
		{
			"November 2021",
			21,
			"    November 2021    ",
		},
		{
			"December 2021",
			24,
			"     December 2021      ",
		},
		{
			"Sol 1990",
			30,
			"           Sol 1990           ",
		},
		{
			"Sol 1990",
			31,
			"           Sol 1990            ",
		},
	} {
		formatted := fmt.CenterInField(input.text, input.fieldWidth)
		if formatted != input.result {
			t.Errorf("%d: Expected '%s' but found '%s'\n", i, input.result, formatted)
		}
	}
}
func TestMonthFormatting(t *testing.T) {
	for i, input := range []struct {
		year   int
		month  cal.IFCMonth
		result []string
	}{
		{
			2021,
			cal.January,
			[]string{
				"      January 2021      ",
				"Su Mo Tu We Th Fr Sa    ",
				" 1  2  3  4  5  6  7    ",
				" 8  9 10 11 12 13 14    ",
				"15 16 17 18 19 20 21    ",
				"22 23 24 25 26 27 28    ",
			},
		},
		{
			2020, // Leap year
			cal.June,
			[]string{
				"       June 2020        ",
				"Su Mo Tu We Th Fr Sa LD ",
				" 1  2  3  4  5  6  7    ",
				" 8  9 10 11 12 13 14    ",
				"15 16 17 18 19 20 21    ",
				"22 23 24 25 26 27 28 29 ",
			},
		},
		{
			2021, // Not a leap year
			cal.June,
			[]string{
				"       June 2021        ",
				"Su Mo Tu We Th Fr Sa    ",
				" 1  2  3  4  5  6  7    ",
				" 8  9 10 11 12 13 14    ",
				"15 16 17 18 19 20 21    ",
				"22 23 24 25 26 27 28    ",
			},
		},
		{
			2021,
			cal.December,
			[]string{
				"     December 2021      ",
				"Su Mo Tu We Th Fr Sa YD ",
				" 1  2  3  4  5  6  7    ",
				" 8  9 10 11 12 13 14    ",
				"15 16 17 18 19 20 21    ",
				"22 23 24 25 26 27 28 29 ",
			},
		},
	} {
		monthFormatting := fmt.MonthToLines(input.year, input.month)
		lineCount := int(math.Min(float64(len(monthFormatting)), float64(len(input.result))))

		for j := 0; j < lineCount; j++ {
			if input.result[j] != monthFormatting[j] {
				t.Errorf("%d: Expected '%s' but found '%s'\n", i, input.result[j], monthFormatting[j])
			}
		}

		if len(input.result) != len(monthFormatting) {
			t.Errorf("%d: Expected %d lines in result but found %d (was: %+v)\n",
				i, len(input.result), len(monthFormatting), monthFormatting)
		}
	}
}
