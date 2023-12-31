package valueobjects

import (
	"fmt"
	"net/http"
	"time"

	shared "rhSystem_server/app/application/error"
)

type Datetime struct {
	Value string
	//Value int64
}

func New(year int, month time.Month, day int, hour int) (*Datetime, *shared.AppError) {

	location, loadLocationErr := time.LoadLocation("America/Sao_Paulo")

	if loadLocationErr != nil {
		return nil, &shared.AppError{Message: "Ocorreu um erro no servidor", StatusCode: http.StatusInternalServerError}
	}

	fmt.Println("year: ", year)
	fmt.Println("month: ", month)
	fmt.Println("day: ", day)

	appointmentDate := time.Date(year, month, day, hour, 0, 0, 0, location)

	//fmt.Println("appointmentDate: ", time.Unix(appointmentDate.Unix(), 0).String())
	fmt.Println("appointmentDate: ", time.Unix(appointmentDate.Unix(), 0).Format("2006-01-02 15:04:05"))

	return &Datetime{
		Value: time.Unix(appointmentDate.Unix(), 0).Format("2006-01-02 15:04:05"),
		//Value: appointmentDate.Unix(),
	}, nil
}

// criar um outro usuário e tentar marcar no mesmo horário -> não pode passar no teste
