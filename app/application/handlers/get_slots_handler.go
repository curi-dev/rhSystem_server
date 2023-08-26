package handlers

import (
	"encoding/json"
	"net/http"
	"rhSystem_server/app/domain/appointments/services"
)

func GetSlotsHandler(w http.ResponseWriter, r *http.Request) {

	// helpers.EnableCors(&w)

	slots, err := services.GetSlotsService()

	if err != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusInternalServerError)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(slots)

	if encodeErr != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
