package services

import (
	shared "rhSystem_server/app/application/error"
	repositories "rhSystem_server/app/infrastructure/repositories/appointments"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func ConfirmAppointmentService(id string) (bool, *shared.AppError) {
	var repo interfaces.AppointmentsRepositoryInterface
	repo = repositories.New()

	return repo.UpdateStatusToConfirmed(id)
}
