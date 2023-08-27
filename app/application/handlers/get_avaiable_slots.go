package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/usecases"
)

func GetAvaiableSlotsHandler(w http.ResponseWriter, r *http.Request) {

	// helpers.EnableCors(&w)

	var avaiableSlotsDTO dtos.AvaiableSlotsRequestDTO

	//The result from u.Query() is of type url.Values, which is a map of strings to slices of strings.
	//Each key in the map corresponds to a query parameter name,
	//and the associated value is a slice of strings containing all the values for that query parameter.
	fmt.Println("r.URL.RawQuery: ", r.URL.RawQuery)
	u, parseErr := url.ParseQuery(r.URL.RawQuery)
	if parseErr != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusBadRequest)
		return
	}

	fmt.Println("u: ", u)

	avaiableSlotsDTO.Day = u.Get("day")
	avaiableSlotsDTO.Month = u.Get("month")
	avaiableSlotsDTO.Year = u.Get("year")

	fmt.Println("avaiableSlotsDTO: ", avaiableSlotsDTO)

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
