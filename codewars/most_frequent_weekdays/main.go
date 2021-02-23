package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	currentYear := 1167
	currentLocation := now.Location()

	firstOfYear := time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentLocation)
	lastOfYear := firstOfYear.AddDate(0, 12, 0)
	numberOfDays := lastOfYear.Sub(firstOfYear).Hours() / 24

	var listOfDays []string
	for i := 0; i < int(numberOfDays); i++ {

		nextDay := firstOfYear.AddDate(0, 0, i)
		listOfDays = append(listOfDays, nextDay.Weekday().String())

	}

	dict := make(map[string]int)
	for _, day := range listOfDays {
		dict[day] = dict[day] + 1
	}

	fmt.Println(dict)

}
