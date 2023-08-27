package usecases

import (
	"fmt"
	"net/http"
	shared "rhSystem_server/app/application/error"
	appointmentsServices "rhSystem_server/app/domain/appointments/services"
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
		return nil, &shared.AppError{Err: nil, Message: "Candidato não encontrado", StatusCode: http.StatusNotFound}
	}

	go func() {
		accessKey, err := candidatesServices.CreateAccessKeyService(candidate.Id.String(), candidatesRepo)

		if err != nil {
			fmt.Println("Error: ", err.Message)
		}

		success := appointmentsServices.SendConfirmationEmail(candidate.Email, accessKey.Value)

		if success {
			fmt.Println("Email sent!")
		}

	}()

	return candidate, nil
}
