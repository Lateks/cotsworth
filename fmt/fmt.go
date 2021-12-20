package fmt

import (
	"fmt"
	"github.com/Lateks/cotsworth/cal"
	"time"
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

func weekdayHeader(year int, month cal.IFCMonth, weekInMonth int) string {
	var eighthDay string
	if weekInMonth == cal.WeeksInMonth {
		if month == cal.December {
			eighthDay = fmt.Sprintf("%s ", cal.YearDay.ShortFormat())
		} else if month == cal.June && cal.IsLeapYear(year) {
			eighthDay = fmt.Sprintf("%s ", cal.LeapDay.ShortFormat())
		} else {
			eighthDay = "   "
		}
	}

	return fmt.Sprintf("%s %s %s %s %s %s %s ",
		cal.Sunday.ShortFormat(), cal.Monday.ShortFormat(), cal.Tuesday.ShortFormat(),
		cal.Wednesday.ShortFormat(), cal.Thursday.ShortFormat(), cal.Friday.ShortFormat(),
		cal.Saturday.ShortFormat()) + eighthDay

}

func formatDay(year int, month cal.IFCMonth, dayNum int, dayToHighlight *cal.IFCDate) string {
	if dayToHighlight != nil && dayToHighlight.Day == dayNum && dayToHighlight.Month == month && dayToHighlight.Year == year {
		return fmt.Sprintf("\033[7m%2d\033[0m ", dayNum)
	}
	return fmt.Sprintf("%2d ", dayNum)
}

func weekLine(weekNum int, year int, month cal.IFCMonth, highlightDay *cal.IFCDate, equalWidth bool) string {
	var eighthDay string
	if weekNum == cal.WeeksInMonth-1 && (month == cal.December || month == cal.June && cal.IsLeapYear(year)) {
		eighthDay = formatDay(year, month, 29, highlightDay)
	} else if equalWidth {
		eighthDay = "   "
	}

	startDay := weekNum*7 + 1
	result := formatDay(year, month, startDay, highlightDay)
	for day := startDay + 1; day < startDay+cal.DaysInWeek; day++ {
		result += formatDay(year, month, day, highlightDay)
	}
	result += eighthDay

	return result
}

func MonthToLines(year int, month cal.IFCMonth, currentDate *cal.IFCDate) []string {
	lines := make([]string, 7)
	title := fmt.Sprintf("%s %d", month, year)
	lines[0] = CenterInField(title, monthWidth)
	lines[1] = weekdayHeader(year, month, cal.WeeksInMonth)

	for week := 0; week < cal.WeeksInMonth; week++ {
		lines[week+2] = weekLine(week, year, month, currentDate, true)
	}

	lines[6] = fmt.Sprintf("%*s", monthWidth, "")

	return lines
}

func formatGregorianChangeOfMonthLine(month time.Month, changeCellIndex int, daysInIFCMonth int) string {
	monthName := month.String()
	charsInMonthName := utf8.RuneCountInString(monthName)
	leftPad := cellWidth * changeCellIndex
	rightPad := cellWidth*daysInIFCMonth - leftPad
	return fmt.Sprintf("%*s%*s", leftPad+charsInMonthName, monthName, rightPad-charsInMonthName, "")
}

func gregorianMonthToLines(startDate time.Time, numDays int, highlightDate time.Time) (dayNumbers string, weekdays string, monthLine string) {
	date := startDate
	for i := 0; i < numDays; i++ {
		dayFormat := "%2d "
		if highlightDate.Equal(date) {
			dayFormat = "\033[7m%2d\033[0m "
		}

		dayNumbers += fmt.Sprintf(dayFormat, date.Day())
		weekdays += fmt.Sprintf("%s ", date.Weekday().String()[:2])
		if date.Day() == 1 {
			monthLine = formatGregorianChangeOfMonthLine(date.Month(), i, numDays)
		}
		date = date.Add(24 * time.Hour)
	}

	return
}

func MonthToLinesWithGregorian(year int, month cal.IFCMonth, currentDate *cal.IFCDate) []string {
	title := fmt.Sprintf("%s %d", month, year)
	weekdays := ""
	dayNumbers := ""
	for i := 0; i < cal.WeeksInMonth; i++ {
		weekdays += weekdayHeader(year, month, i+1)
		dayNumbers += weekLine(i, year, month, currentDate, false)
	}

	var gregorianHighlightDay time.Time
	if currentDate != nil {
		gregorianHighlightDay = currentDate.ToUTCTime()
	}

	gregorianDayNumbers, gregorianWeekdays, gregorianMonthLine := gregorianMonthToLines(
		cal.NewIFCDate(year, month, 1).ToUTCTime(),
		cal.DaysInMonth(year, month),
		gregorianHighlightDay,
	)

	return []string{
		title,
		weekdays,
		dayNumbers,
		gregorianDayNumbers,
		gregorianWeekdays,
		gregorianMonthLine,
	}
}
