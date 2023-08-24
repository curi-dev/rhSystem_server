package services

import (
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func CheckIfSlotAvaiable(slot int, repo interfaces.SlotsRepositoryInterface) (bool, *shared.AppError) {
	return repo.FindById(slot)
}
