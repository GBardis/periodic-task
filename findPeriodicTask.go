package main

import (
	"codeberg.org/eviedelta/detctime/durationparser"
	"github.com/itchyny/timefmt-go"
	"strings"
	"time"
)

func findPeriodicTasks(period, timezone, timeStart, timeEnd string) []string {
	randomDates := randate(timezone)

	start := parseDate(timeStart, timezone)
	end := parseDate(timeEnd, timezone)

	return rangeDate(start, end, randomDates, period)
}

func rangeDate(start, end time.Time, ptlist []time.Time, period string) []string {
	var datesInBetween []string
	intervalDate := start
	for _, d := range ptlist {
		if d.After(start) && d.Before(end) {
			datesInBetween = dateCompare(d, intervalDate, period, datesInBetween)
			intervalDate = periodInterval(intervalDate, period)
		}
	}
	return datesInBetween
}

func dateCompare(ptlistDate, intervalDate time.Time, period string, datesInBetween []string) []string {
	switch {
	case strings.Contains(period, "h"):
		if ptlistDate.Hour() == intervalDate.Hour() {
			date := timefmt.Format(ptlistDate, "%Y%m%dT%H%M%SZ")
			datesInBetween = append(datesInBetween, date)
		}
	case strings.Contains(period, "d"):
		if ptlistDate.Day() == intervalDate.Day() {
			date := timefmt.Format(ptlistDate, "%Y%m%dT%H%M%SZ")
			datesInBetween = append(datesInBetween, date)
		}
	case strings.Contains(period, "m"):
		if ptlistDate.Month().String() == intervalDate.Month().String() {
			date := timefmt.Format(ptlistDate, "%Y%m%dT%H%M%SZ")
			datesInBetween = append(datesInBetween, date)
		}
	case strings.Contains(period, "y"):
		if ptlistDate.Year() == intervalDate.Year() {
			date := timefmt.Format(ptlistDate, "%Y%m%dT%H%M%SZ")
			datesInBetween = append(datesInBetween, date)
		}
	default:
		//fmt.Println("Not valid period param")
	}
	return datesInBetween
}

func periodInterval(startDate time.Time, stringDuration string) time.Time {
	duration, _ := durationparser.Parse(stringDuration)
	return startDate.Add(duration)
}

func parseDate(date string, timezone string) time.Time {
	location, _ := time.LoadLocation(timezone)
	parsedDate, _ := timefmt.Parse(date, "%Y%m%dT%H%M%SZ")
	return parsedDate.UTC().In(location)
}
