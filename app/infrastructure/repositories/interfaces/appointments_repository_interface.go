package interfaces

import (
	shared "rhSystem_server/app/application/error"
)

type AppointmentsRepositoryInterface interface {
	// FindByDatetime(date string, slot int) (*entities.Appointment, *shared.AppError)
	FindByCandidateEmail(email string) (map[string]interface{}, *shared.AppError)
	UpdateStatus(id int, status int) (bool, *shared.AppError)
}
