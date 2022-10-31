package main

import (
	"math/rand"
	"time"
)

func randate(timezone string) []time.Time {
	location, _ := time.LoadLocation(timezone)
	var dates []time.Time

	sum := 0
	for i := 1; i < 1000; i++ {
		sum += i
		min := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		max := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		delta := max - min

		sec := rand.Int63n(delta) + min
		date := time.Unix(sec, 0).UTC().In(location)
		dates = append(dates, date)

	}
	return dates
}
