package dtos

import "rhSystem_server/app/infrastructure/repositories/interfaces"

type UpdateAppointmentStatusDTO struct {
	Id     int
	Status int
	Repo   interfaces.AppointmentsRepositoryInterface
	C      chan bool
}
