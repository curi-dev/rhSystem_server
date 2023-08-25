package interfaces

import (
	shared "rhSystem_server/app/application/error"
	"time"
)

type SlotsRepositoryInterface interface {
	//Index()
	Find(w *time.Weekday, slot int) (bool, *shared.AppError)
}
