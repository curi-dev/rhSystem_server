package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

//func CheckIfCandidateHasAppointmentAlready(id int, repo interfaces.AppointmentsRepositoryInterface) (bool, *shared.AppError) {
func CheckIfCandidateHasAppointmentAlready(email string, repo interfaces.AppointmentsRepositoryInterface) (map[string]interface{}, *shared.AppError) {
	return repo.FindByCandidateEmail(email)
}
