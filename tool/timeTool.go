package timetool

import (
	"log"
	"time"
)

func ParseDate(dateString string) time.Time {
	date, err := time.Parse("02-01-2006", dateString)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Parsing date", date)

	return date
}
