package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func UpdateAppointmentStatusService(id int, status int, repo interfaces.AppointmentsRepositoryInterface) (bool, *shared.AppError) {
	return repo.UpdateStatus(id, status)
}
