package interfaces

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/entities"
)

type AppointmentsRepositoryInterface interface {
	FindByCandidateId(candidateId string) (map[string]interface{}, *shared.AppError)
	UpdateStatus(id int, status int) (bool, *shared.AppError)
	Create(a *entities.Appointment, candidateId string) (bool, *shared.AppError)
	FindByDatetime(datetime string) ([]int, *shared.AppError)
}
