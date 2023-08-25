package services

import (
	"fmt"
	"rhSystem_server/app/domain/appointments/dtos"
)

func UpdateAppointmentStatusOnRoutineService(updateAppointmentDTO dtos.UpdateAppointmentStatusDTO) {
	go func() {
		proceed := <-updateAppointmentDTO.C

		fmt.Println("proceed: ", proceed)

		if proceed {
			updateAppointmentDTO.Repo.UpdateStatus(updateAppointmentDTO.Id, updateAppointmentDTO.Status)
		}
	}()
}
