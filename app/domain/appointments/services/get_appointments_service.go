package services

import (
	shared "rhSystem_server/app/application/error"
	repositories "rhSystem_server/app/infrastructure/repositories/appointments"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func GetAppointmentsService() ([]interface{}, *shared.AppError) {

	var repo interfaces.AppointmentsRepositoryInterface
	repo = repositories.New()

	resp, err := repo.Index()

	if err != nil {
		return nil, err
	}

	return resp, nil
}
