package services

import (
	shared "rhSystem_server/app/application/error"
	aggregate "rhSystem_server/app/domain/appointments/aggregate/validSlot"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
	"time"
)

func GetWeekDayValidSlotsService(weekdayId time.Weekday, repo interfaces.SlotsRepositoryInterface) ([]aggregate.ValidSlot, *shared.AppError) {
	return repo.FindByWeekday(int(weekdayId))
}
