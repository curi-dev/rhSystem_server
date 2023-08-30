package interfaces

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/entities"

	"github.com/google/uuid"
)

type AppointmentsRepositoryInterface interface {
	// defaults
	Index() ([]interface{}, *shared.AppError)
	Create(a *entities.Appointment, candidateId string) (bool, *shared.AppError)

	// specific defaults
	FindByCandidateId(candidateId uuid.UUID) (map[string]interface{}, *shared.AppError)
	FindByDatetime(datetime string) ([]int, *shared.AppError)
	UpdateStatus(id int, status int) (bool, *shared.AppError)

	// specific query based
	FindBlockedSlotsByDatetime(datetime string) ([]int, *shared.AppError)
	UpdateStatusToConfirmed(id string) (bool, *shared.AppError)
}
