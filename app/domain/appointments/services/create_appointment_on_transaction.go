package services

import (
	"rhSystem_server/app/domain/appointments/entities"

	shared "rhSystem_server/app/application/error"
	repositories "rhSystem_server/app/infrastructure/repositories/transactions"
)

func CreateAppointmentOnTransactionService(candidate *entities.Candidate, appointment *entities.Appointment) (bool, *shared.AppError) {
	repo := repositories.New() // convert into interface like the others
	return repo.Run(candidate, appointment)
}
