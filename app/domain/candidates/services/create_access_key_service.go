package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/helpers"
	valueobjects "rhSystem_server/app/domain/candidates/valueobjects/accessKey"
	"rhSystem_server/app/infrastructure/repositories/interfaces"

	"github.com/google/uuid"
)

func CreateAccessKeyService(candidateId string, repo interfaces.CandidateRepositoryInterface) (*valueobjects.AccessKey, *shared.AppError) {
	randomKey := helpers.GenerateRandomKey(8)

	accessKey := valueobjects.AccessKey{Id: uuid.New(), Value: randomKey, Candidate: candidateId}

	return repo.AccessKey(&accessKey)
}
