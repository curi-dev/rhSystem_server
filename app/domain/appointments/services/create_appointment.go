package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/entities"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func CreateAppointmentService(newAppointment *entities.Appointment, candidateId string, repo interfaces.AppointmentsRepositoryInterface) (bool, *shared.AppError) {
	return repo.Create(newAppointment, candidateId)
}
