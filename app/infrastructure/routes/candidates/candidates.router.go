package routes

import (
	"net/http"
	"rhSystem_server/app/application/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// v1
	mux.HandleFunc("/api/v1/candidates/create", func(w http.ResponseWriter, r *http.Request) { handlers.CreateCandidateHandler(w, r) })

	return mux
}
