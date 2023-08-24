package routes

import (
	"net/http"
	// "rhSystem_server/app/application/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// v1
	mux.HandleFunc("/api/v1/appointments", func(w http.ResponseWriter, r *http.Request) {})
	//mux.HandleFunc("/api/v1/companies/suggest", func(w http.ResponseWriter, r *http.Request) { handlers.SearchCompanyHandler(w, r) })

	return mux
}
