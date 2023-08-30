package interfaces

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/entities"
	valueobjects "rhSystem_server/app/domain/candidates/valueobjects/accessKey"
)

type CandidateRepositoryInterface interface {
	// candidate
	Create(c *entities.Candidate) (*entities.Candidate, *shared.AppError)

	FindByEmail(email string) (*entities.Candidate, *shared.AppError)
	FindById(id string) (*entities.Candidate, *shared.AppError)

	// access key
	AccessKey(k *valueobjects.AccessKey) (*valueobjects.AccessKey, *shared.AppError)
	FindKeyByCandidateId(candidateId string) (*valueobjects.AccessKey, *shared.AppError)
}
