package fmt

import (
	"fmt"
	"github.com/Lateks/cotsworth/cal"
	"unicode/utf8"
)

const (
	cellWidth      = 3
	daysOnWeekLine = 8
	monthWidth     = daysOnWeekLine * cellWidth
)

func CenterInField(text string, fieldWidth int) string {
	textWidth := utf8.RuneCountInString(text)
	totalPad := fieldWidth - textWidth
	leftPad := totalPad / 2
	return fmt.Sprintf("%-*s", fieldWidth, fmt.Sprintf("%*s", leftPad+textWidth, text))
}

func weekdayHeader(year int, month cal.IFCMonth) string {
	eighthDay := "  "
	if month == cal.December {
		eighthDay = cal.YearDay.ShortFormat()
	} else if month == cal.June && cal.IsLeapYear(year) {
		eighthDay = cal.LeapDay.ShortFormat()
	}

	return fmt.Sprintf("%s %s %s %s %s %s %s %s ",
		cal.Sunday.ShortFormat(), cal.Monday.ShortFormat(), cal.Tuesday.ShortFormat(),
		cal.Wednesday.ShortFormat(), cal.Thursday.ShortFormat(), cal.Friday.ShortFormat(),
		cal.Saturday.ShortFormat(), eighthDay)
}

func dayLine(weekNum int, year int, month cal.IFCMonth) string {
	var eighthDay string
	if weekNum == cal.WeeksInMonth-1 && (month == cal.December || month == cal.June && cal.IsLeapYear(year)) {
		eighthDay = "29"
	} else {
		eighthDay = "  "
	}

	startDay := weekNum*7 + 1
	return fmt.Sprintf("%2d %2d %2d %2d %2d %2d %2d %s ",
		startDay, startDay+1, startDay+2, startDay+3, startDay+4, startDay+5, startDay+6, eighthDay)
}

func MonthToLines(year int, month cal.IFCMonth) []string {
	lines := make([]string, 6)
	title := fmt.Sprintf("%s %d", month, year)
	lines[0] = CenterInField(title, monthWidth)
	lines[1] = weekdayHeader(year, month)

	for week := 0; week < cal.WeeksInMonth; week++ {
		lines[week+2] = dayLine(week, year, month)
	}

	return lines
}
