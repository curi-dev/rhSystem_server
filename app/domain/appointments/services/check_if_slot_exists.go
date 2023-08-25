package services

import (
	"net/http"
	"time"

	shared "rhSystem_server/app/application/error"
	weekday "rhSystem_server/app/domain/appointments/valueobjects"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func CheckIfSlotExists(slot int, splittedDate *struct { // what is going on here?
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}, repo interfaces.SlotsRepositoryInterface) (bool, *shared.AppError) {
	location, loadLocationErr := time.LoadLocation("America/Sao_Paulo")

	if loadLocationErr != nil {
		return false, &shared.AppError{Message: "Ocorreu um erro no servidor", StatusCode: http.StatusInternalServerError}
	}

	w := weekday.New(splittedDate.Year, splittedDate.Month, splittedDate.Day, location)

	return repo.Find(&w.Value, slot) // what is actually passed here?
}
