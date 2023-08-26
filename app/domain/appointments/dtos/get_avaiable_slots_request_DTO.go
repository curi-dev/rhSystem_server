package dtos

import (
	"time"
)

type AvaiableSlotsRequestDTO struct {
	SplittedDate struct {
		Year  int        `json:"year"`
		Month time.Month `json:"month"`
		Day   int        `json:"day"`
	} `json:"splitted_date"`
}
