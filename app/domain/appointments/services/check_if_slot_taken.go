package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/entities"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func CheckIfSlotTaken(date string, slot int, repo interfaces.AppointmentsRepositoryInterface) (*entities.Appointment, *shared.AppError) {
	return repo.FindByDatetime(date, slot)
}
