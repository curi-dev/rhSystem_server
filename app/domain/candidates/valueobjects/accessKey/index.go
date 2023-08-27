package valueobjects

import "github.com/google/uuid"

type AccessKey struct {
	Id        uuid.UUID
	Value     string
	Candidate string
}
