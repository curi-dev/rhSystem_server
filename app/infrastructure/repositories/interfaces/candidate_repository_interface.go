package interfaces

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/entities"
)

type CandidateRepositoryInterface interface {
	FindByEmail(email string) (bool, *shared.AppError)
	Create(c *entities.Candidate) (*entities.Candidate, *shared.AppError)
}
