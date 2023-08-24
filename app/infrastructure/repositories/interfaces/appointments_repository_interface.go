package interfaces

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/entities"
)

type AppointmentsRepositoryInterface interface {
	FindByDatetime(date string, slot int) (*entities.Appointment, *shared.AppError)
}
