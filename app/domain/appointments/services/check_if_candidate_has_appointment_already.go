package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

//func CheckIfCandidateHasAppointmentAlready(id int, repo interfaces.AppointmentsRepositoryInterface) (bool, *shared.AppError) {
func CheckIfCandidateHasAppointmentAlready(id string, repo interfaces.AppointmentsRepositoryInterface) (map[string]interface{}, *shared.AppError) {
	return repo.FindByCandidateId(id)
}
