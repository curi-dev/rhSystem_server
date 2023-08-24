package interfaces

import shared "rhSystem_server/app/application/error"

type CandidateRepositoryInterface interface {
	FindByEmail(email string) (bool, *shared.AppError)
}
