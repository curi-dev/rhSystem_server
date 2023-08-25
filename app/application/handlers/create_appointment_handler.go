package handlers

import (
	"encoding/json"
	"net/http"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/usecases"
)

func CreateAppointmentHandler(w http.ResponseWriter, r *http.Request) {

	var newAppointment dtos.NewAppointmentRequestDTO

	decodeErr := json.NewDecoder(r.Body).Decode(&newAppointment)

	if decodeErr != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusInternalServerError)
		return
	}

	_, err := usecases.CreateAppointmentUseCase(&newAppointment)

	if err != nil {
		http.Error(w, err.Err.Error(), err.StatusCode)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
