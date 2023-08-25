package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/dtos"
	"rhSystem_server/app/domain/candidates/entities"
	"rhSystem_server/app/infrastructure/repositories/interfaces"

	"github.com/google/uuid"
)

func CreateCandidateService(newCandidateDTO *dtos.CandidateRequestDTO, repo interfaces.CandidateRepositoryInterface) (*entities.Candidate, *shared.AppError) {

	newCandidate := entities.Candidate{Id: uuid.New(), Name: newCandidateDTO.Name, Email: newCandidateDTO.Email, Phone: newCandidateDTO.Phone}

	return repo.Create(&newCandidate)
}
