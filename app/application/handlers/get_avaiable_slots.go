package handlers

import (
	"encoding/json"
	"net/http"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/usecases"
)

func GetAvaiableSlotsHandler(w http.ResponseWriter, r *http.Request) {

	// helpers.EnableCors(&w)

	var avaiableSlotsDTO dtos.AvaiableSlotsRequestDTO

	decoderErr := json.NewDecoder(r.Body).Decode(&avaiableSlotsDTO)

	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	avaiableSlots, err := usecases.GetAvaiableSlotsUseCase(&avaiableSlotsDTO)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	encodeErr := json.NewEncoder(w).Encode(avaiableSlots)

	if encodeErr != nil {
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		return
	}

}
