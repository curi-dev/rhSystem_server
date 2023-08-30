package services

import (
	"net/http"
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/entities"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func FindCandidateById(id string, repo interfaces.CandidateRepositoryInterface) (*entities.Candidate, *shared.AppError) {
	candidate, err := repo.FindById(id)

	if err != nil {
		return nil, err
	}

	if candidate == nil {
		return nil, &shared.AppError{Err: nil, Message: "Candidato não encontrado", StatusCode: http.StatusNotFound}
	}

	return candidate, nil
}
