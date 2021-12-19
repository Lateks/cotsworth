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

func formatDay(dayNum int, dayToHighlight int) string {
	if dayNum == dayToHighlight {
		return fmt.Sprintf("\033[7m%2d\033[0m ", dayNum)
	}
	return fmt.Sprintf("%2d ", dayNum)
}

func dayLine(weekNum int, year int, month cal.IFCMonth, highlightDay int) string {
	var eighthDay string
	if weekNum == cal.WeeksInMonth-1 && (month == cal.December || month == cal.June && cal.IsLeapYear(year)) {
		eighthDay = formatDay(29, highlightDay)
	} else {
		eighthDay = "   "
	}

	startDay := weekNum*7 + 1
	result := formatDay(startDay, highlightDay)
	for day := startDay + 1; day < startDay+cal.DaysInWeek; day++ {
		result += formatDay(day, highlightDay)
	}
	result += eighthDay

	return result
}

func MonthToLines(year int, month cal.IFCMonth, currentDate *cal.IFCDate) []string {
	lines := make([]string, 7)
	title := fmt.Sprintf("%s %d", month, year)
	lines[0] = CenterInField(title, monthWidth)
	lines[1] = weekdayHeader(year, month)

	highlightDay := -1
	if currentDate != nil && currentDate.Year == year && currentDate.Month == month {
		highlightDay = currentDate.Day
	}

	for week := 0; week < cal.WeeksInMonth; week++ {
		lines[week+2] = dayLine(week, year, month, highlightDay)
	}

	lines[6] = fmt.Sprintf("%*s", monthWidth, "")

	return lines
}
