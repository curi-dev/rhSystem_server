package dtos

import (
	"time"

	"github.com/google/uuid"
)

type AppointmentRequestDTO struct {
	CandidateId    uuid.UUID `json:"id"`
	CandidateEmail string    `json:"email"`
	// Datetime       string    `json:"datetime"`
	Slot         int `json:"slot"`
	SplittedDate struct {
		Year  int        `json:"year"`
		Month time.Month `json:"month"`
		Day   int        `json:"day"`
	} `json:"splitted_date"`
}
