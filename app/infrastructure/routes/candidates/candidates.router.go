package routes

import (
	"net/http"
	"rhSystem_server/app/application/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// v1
	mux.HandleFunc("/api/v1/candidates/create", func(w http.ResponseWriter, r *http.Request) { handlers.CreateCandidateHandler(w, r) })
	mux.HandleFunc("/api/v1/candidates/access-key/create", func(w http.ResponseWriter, r *http.Request) { handlers.CreateAccessKeyHandler(w, r) })
	mux.HandleFunc("/api/v1/candidates/access-key/use", func(w http.ResponseWriter, r *http.Request) { handlers.UseKeyAccessHandler(w, r) })

	return mux
}
