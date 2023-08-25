package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/entities"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func CreateCandidateService(newCandidate *entities.Candidate, repo interfaces.CandidateRepositoryInterface) (*entities.Candidate, *shared.AppError) {
	return repo.Create(newCandidate)
}
