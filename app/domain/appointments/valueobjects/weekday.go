package valueobjects

import (
	"fmt"
	"time"
)

type Weekday struct {
	//Name  string
	Value time.Weekday
}

func New(year int, month time.Month, day int, location *time.Location) *Weekday {

	appointmentDate := time.Date(year, month, day, 0, 0, 0, 0, location)

	fmt.Println("appointmentDate: ", appointmentDate)

	return &Weekday{
		Value: appointmentDate.Weekday(),
	}
}
