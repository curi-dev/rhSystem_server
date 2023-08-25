package dtos

import (
	"time"

	"github.com/google/uuid"
)

type NewAppointmentRequestDTO struct {
	CandidateId  uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Datetime     string    `json:"datetime"`
	Slot         int       `json:"slot"`
	SplittedDate struct {
		Year  int        `json:"year"`
		Month time.Month `json:"month"`
		Day   int        `json:"day"`
	} `json:"splited_date"`
}
