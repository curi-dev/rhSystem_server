package interfaces

import (
	shared "rhSystem_server/app/application/error"
	slotAggregate "rhSystem_server/app/domain/appointments/aggregate/slot"
	validSlotAggregate "rhSystem_server/app/domain/appointments/aggregate/validSlot"
	"time"
)

type SlotsRepositoryInterface interface {
	Index() ([]slotAggregate.Slot, *shared.AppError)
	Find(w *time.Weekday, slot int) (bool, *shared.AppError)
	FindByWeekday(weekdayId int) ([]validSlotAggregate.ValidSlot, *shared.AppError)
}
