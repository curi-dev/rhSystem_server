package services

import (
	"fmt"
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

//func FilterAppointmentsByDatetimeService(datetime string, repo interfaces.AppointmentsRepositoryInterface) ([]int, *shared.AppError) {
func FindBlockedSlotsByDatetimeService(datetime string, repo interfaces.AppointmentsRepositoryInterface) ([]int, *shared.AppError) {

	fmt.Println("datetime: [FilterAppointmentsByDatetimeService]", datetime)

	return repo.FindBlockedSlotsByDatetime(datetime)
}
