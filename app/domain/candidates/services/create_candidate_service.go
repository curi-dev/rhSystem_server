package services

import (
	"net/http"

	"github.com/google/uuid"

	"rhSystem_server/app/domain/candidates/dtos"
	"rhSystem_server/app/domain/candidates/entities"
	valueobjects "rhSystem_server/app/domain/candidates/valueobjects/email"
	"rhSystem_server/app/infrastructure/repositories/interfaces"

	shared "rhSystem_server/app/application/error"
)

func CreateCandidateService(newCandidateDTO *dtos.CandidateRequestDTO, repo interfaces.CandidateRepositoryInterface) (*entities.Candidate, *shared.AppError) {

	email := valueobjects.New(newCandidateDTO.Email)

	if !email.IsValid() ||
		len(newCandidateDTO.Name) < 1 ||
		len(newCandidateDTO.Phone) < 11 {
		return nil, &shared.AppError{Message: "Dados inválidos", StatusCode: http.StatusBadRequest}
	}

	newCandidate := entities.Candidate{Id: uuid.New(), Name: newCandidateDTO.Name, Email: newCandidateDTO.Email, Phone: newCandidateDTO.Phone}

	return repo.Create(&newCandidate)
}
