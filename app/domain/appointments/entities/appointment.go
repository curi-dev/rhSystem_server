package entities

import "github.com/google/uuid"

type Appointment struct {
	Id        uuid.UUID
	Datetime  string
	Slot      int
	Candidate uuid.UUID // Candidate
	Status    int       // 'suspense' | 'confirmed' -> Criar enum
	ResumeUrl string
}
