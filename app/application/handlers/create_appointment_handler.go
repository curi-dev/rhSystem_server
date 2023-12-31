package handlers

import (
	"encoding/json"
	"net/http"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/usecases"
)

func CreateAppointmentHandler(w http.ResponseWriter, r *http.Request) {

	var newAppointmentDTO dtos.AppointmentRequestDTO

	// como o formato DATE do banco se comporta com o datetime que vem do json
	decodeErr := json.NewDecoder(r.Body).Decode(&newAppointmentDTO)

	if decodeErr != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	success, err := usecases.CreateAppointmentUseCase(&newAppointmentDTO)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	if !success {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	//w.Header().Set("Content-Type", "application/json")
}
