package services

import (
	"net/http"
	shared "rhSystem_server/app/application/error"
	valueobjects "rhSystem_server/app/domain/candidates/valueobjects/accessKey"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func FindKeyService(column string, value string, repo interfaces.CandidateRepositoryInterface) (*valueobjects.AccessKey, *shared.AppError) {
	k, err := repo.FindKeyByCandidateId(value)

	if err != nil {
		return nil, err
	}

	if k == nil {
		return nil, &shared.AppError{Err: nil, Message: "Chave de acesso inválida", StatusCode: http.StatusBadRequest}
	}

	if k.IsValid() {
		return k, nil
	}

	return nil, &shared.AppError{Err: nil, Message: "Chave de acesso inválida", StatusCode: http.StatusBadRequest}
}
