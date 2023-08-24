package dtos

type NewAppointmentRequestDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Datetime string `json:"datetime"`
	Slot     int    `json:"slot"`
}
