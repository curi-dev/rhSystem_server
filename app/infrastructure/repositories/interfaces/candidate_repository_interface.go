package interfaces

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/entities"
	valueobjects "rhSystem_server/app/domain/candidates/valueobjects/accessKey"
)

type CandidateRepositoryInterface interface {
	FindByEmail(email string) (*entities.Candidate, *shared.AppError)
	Create(c *entities.Candidate) (*entities.Candidate, *shared.AppError)
	AccessKey(k *valueobjects.AccessKey) (*valueobjects.AccessKey, *shared.AppError)
}
