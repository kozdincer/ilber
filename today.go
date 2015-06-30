package main

import (
	"fmt"
	"time"
)

func init() {
	register("/bugun", today)
}

type weekday time.Weekday

var days = [...]string{
	"Pazar",
	"Pazartesi",
	"Sali",
	"Carsamba",
	"Persembe",
	"Cuma",
	"Cumartesi",
}

func (w weekday) String() string {
	return days[w]
}

func today(args ...string) string {
	return fmt.Sprintf("bugun %v", weekday(time.Now().Weekday()).String())
}
