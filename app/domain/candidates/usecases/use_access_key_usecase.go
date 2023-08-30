package usecases

import (
	"net/http"
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/dtos"
	"rhSystem_server/app/domain/candidates/services"
	repositories "rhSystem_server/app/infrastructure/repositories/candidates"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func UseAccessKeyUseCase(usekeyAccessDto *dtos.UseKeyAccessRequestDTO) (bool, *shared.AppError) {

	var candidatesRepo interfaces.CandidateRepositoryInterface
	candidatesRepo = repositories.New()

	// service is responsible for returning a valid key
	k, err := services.FindKeyService("candidate", usekeyAccessDto.Candidate, candidatesRepo)

	if err != nil {
		return false, err
	}

	if k.Value != usekeyAccessDto.Key {
		return false, &shared.AppError{Err: nil, Message: "Chave de acesso inválida", StatusCode: http.StatusBadRequest}
	}

	return true, nil
}
