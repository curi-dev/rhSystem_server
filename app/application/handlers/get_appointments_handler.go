package handlers

import (
	"encoding/json"
	"net/http"
	"rhSystem_server/app/domain/appointments/services"
)

func GetAppointmentsHandler(w http.ResponseWriter, r *http.Request) {

	appointments, err := services.GetAppointmentsService()

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(appointments)

	if encodeErr != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
