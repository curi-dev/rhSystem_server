package routes

import (
	"net/http"
	"rhSystem_server/app/application/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// v1
	mux.HandleFunc("/api/v1/slots/index", func(w http.ResponseWriter, r *http.Request) { handlers.GetSlotsHandler(w, r) })
	mux.HandleFunc("/api/v1/slots/filter", func(w http.ResponseWriter, r *http.Request) { handlers.GetAvaiableSlotsHandler(w, r) })

	return mux
}
