package valueobjects

import (
	"net/http"
	shared "rhSystem_server/app/application/error"
)

type Hour struct {
	Value int
}

func New(slot int) (*Hour, *shared.AppError) {

	var hour int
	switch slot {
	case 1:
		hour = 8
	case 2:
		hour = 9
	case 3:
		hour = 10
	case 4:
		hour = 11
	case 5:
		hour = 12
	case 6:
		hour = 13
	case 7:
		hour = 14
	case 8:
		hour = 15
	case 9:
		hour = 16
	case 10:
		hour = 17
	default:
		return nil, &shared.AppError{Message: "Slot inválido", StatusCode: http.StatusInternalServerError}
	}

	return &Hour{
		Value: hour,
	}, nil
}
