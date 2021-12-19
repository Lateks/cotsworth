package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lateks/cotsworth/cal"
	fcalFmt "github.com/Lateks/cotsworth/fmt"
)

func displayMonth(currentDate *cal.IFCDate) {
	monthLines := fcalFmt.MonthToLines(currentDate.Year, currentDate.Month, currentDate)
	fmt.Println(strings.Join(monthLines, "\n"))
}

func Execute() {
	today := cal.DateAt(time.Now())
	displayMonth(today)
}
