package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/repositories/interfaces"

	"github.com/google/uuid"
)

//func CheckIfCandidateHasAppointmentAlready(id int, repo interfaces.AppointmentsRepositoryInterface) (bool, *shared.AppError) {
func CheckIfCandidateHasAppointmentAlready(id uuid.UUID, repo interfaces.AppointmentsRepositoryInterface) (map[string]interface{}, *shared.AppError) {
	return repo.FindByCandidateId(id)
}
