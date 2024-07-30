// File:		time.go
// Created by:	Hoven
// Created on:	2024-07-30
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package putils

import (
	"fmt"
	"time"
)

func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// FirstDayOfMonth Get the first day of the month in which the given time is located
func FirstDayOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

// LastDayOfMonth Get the last day of the month in which the given time is located
func LastDayOfMonth(date time.Time) time.Time {
	firstDayOfNextMonth := FirstDayOfMonth(date).AddDate(0, 1, 0)
	return firstDayOfNextMonth.Add(-time.Second)
}

// FirstDayOfWeek
func FirstDayOfWeek(date time.Time) time.Time {
	daysSinceSunday := int(date.Weekday())
	return StartOfDay(date).AddDate(0, 0, -daysSinceSunday+1)
}

func LastDayOfWeek(date time.Time) time.Time {
	daysUntilSaturday := 7 - int(date.Weekday())
	return StartOfDay(date).AddDate(0, 0, daysUntilSaturday)
}

func FormatDuration(duration time.Duration) string {
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%d day %02d hour %02d min %02d sec", days, hours, minutes, seconds)
}
