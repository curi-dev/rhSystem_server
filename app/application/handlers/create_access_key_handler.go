package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"rhSystem_server/app/application/helpers"
	"rhSystem_server/app/domain/candidates/usecases"
)

func CreateAccessKeyHandler(w http.ResponseWriter, r *http.Request) {

	helpers.EnableCors(&w)

	fmt.Println("r.URL.RawQuery: ", r.URL.RawQuery)

	u, parseErr := url.ParseQuery(r.URL.RawQuery)
	if parseErr != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusBadRequest)
		return
	}

	var email string
	email = u.Get("email")

	candidate, err := usecases.CreateAccessKeyUseCase(email)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	if candidate != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		encodeErr := json.NewEncoder(w).Encode(candidate)

		if encodeErr != nil {
			http.Error(w, "Ocorreu um problema no servidor", http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}
