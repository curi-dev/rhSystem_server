package usecases

import (
	"fmt"
	shared "rhSystem_server/app/application/error"
	applicationServices "rhSystem_server/app/application/services"
	"rhSystem_server/app/domain/candidates/entities"
	candidatesServices "rhSystem_server/app/domain/candidates/services"
	repositories "rhSystem_server/app/infrastructure/repositories/candidates"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func CreateAccessKeyUseCase(email string) (*entities.Candidate, *shared.AppError) {

	// verify if candidate exists
	var candidatesRepo interfaces.CandidateRepositoryInterface
	candidatesRepo = repositories.New()

	candidate, err := candidatesServices.CheckIfCandidateExistsService(email, candidatesRepo)

	// erro interno [500]
	if err != nil {
		return nil, err
	}

	// candidate does not exist on database (resource not found)
	if candidate == nil {
		//return nil, &shared.AppError{Err: nil, Message: "Candidato não encontrado", StatusCode: http.StatusUnauthorized}
		return nil, nil
	}

	go func() {
		accessKey, err := candidatesServices.CreateAccessKeyService(candidate.Id.String(), candidatesRepo)

		if err != nil {
			fmt.Println("Error: ", err.Message)
		}

		subject := "CHAVE DE ACESSO"
		body := accessKey.Value
		success := applicationServices.SendEmail(candidate.Email, subject, body)

		if success {
			fmt.Println("Email sent!")
		}

	}()

	return candidate, nil
}
