package handlers

import (
	"net/http"
	"net/url"
	"rhSystem_server/app/domain/appointments/services"
)

func ConfirmAppointmentHandler(w http.ResponseWriter, r *http.Request) {

	u, parseErr := url.ParseQuery(r.URL.RawQuery)
	if parseErr != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusBadRequest)
		return
	}

	appointmentId := u.Get("appointment")

	success, err := services.ConfirmAppointmentService(appointmentId)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	if !success {
		http.Error(w, err.Message, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
