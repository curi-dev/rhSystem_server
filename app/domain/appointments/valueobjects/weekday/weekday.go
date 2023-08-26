package valueobjects

import (
	"fmt"
	"net/http"
	shared "rhSystem_server/app/application/error"
	"time"
)

type Weekday struct {
	//Name  string
	Value time.Weekday
}

func New(year int, month time.Month, day int) (*Weekday, *shared.AppError) {
	location, loadLocationErr := time.LoadLocation("America/Sao_Paulo")

	if loadLocationErr != nil {
		return nil, &shared.AppError{Message: "Ocorreu um erro no servidor", StatusCode: http.StatusInternalServerError}
	}

	appointmentDate := time.Date(year, month, day, 0, 0, 0, 0, location)

	fmt.Println("appointmentDate [weekday]: ", appointmentDate.Weekday())
	fmt.Println("appointmentDate: ", appointmentDate)

	return &Weekday{
		Value: appointmentDate.Weekday(),
	}, nil
}
