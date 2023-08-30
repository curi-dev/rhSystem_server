package routes

import (
	"net/http"
	"rhSystem_server/app/application/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// v1
	mux.HandleFunc("/api/v1/appointments/create", func(w http.ResponseWriter, r *http.Request) { handlers.CreateAppointmentHandler(w, r) })
	mux.HandleFunc("/api/v1/appointments/index", func(w http.ResponseWriter, r *http.Request) { handlers.GetAppointmentsHandler(w, r) })
	mux.HandleFunc("/api/v1/appointments/confirm", func(w http.ResponseWriter, r *http.Request) { handlers.ConfirmAppointmentHandler(w, r) })

	return mux
}
