package entities

import "github.com/google/uuid"

type Candidate struct {
	Id        uuid.UUID
	Name      string
	Email     string
	Phone     string
	ResumeUrl string
}
