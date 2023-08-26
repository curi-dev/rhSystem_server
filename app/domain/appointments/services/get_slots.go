package services

import (
	shared "rhSystem_server/app/application/error"
	aggregate "rhSystem_server/app/domain/appointments/aggregate/slot"
	repositories "rhSystem_server/app/infrastructure/repositories/slots"
)

func GetSlotsService() ([]aggregate.Slot, *shared.AppError) {
	return repositories.New().Index()
}
