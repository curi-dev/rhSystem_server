package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"rhSystem_server/app/domain/candidates/dtos"
	"rhSystem_server/app/domain/candidates/usecases"
)

func UseKeyAccessHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("r.URL.RawQuery: ", r.URL.RawQuery)

	u, parseErr := url.ParseQuery(r.URL.RawQuery)
	if parseErr != nil {
		http.Error(w, "Ocorreu um problema no servidor", http.StatusBadRequest)
		return
	}

	var accessKeyRequestDto dtos.UseKeyAccessRequestDTO
	accessKeyRequestDto.Candidate = u.Get("candidate")
	accessKeyRequestDto.Key = u.Get("key")

	success, err := usecases.UseAccessKeyUseCase(&accessKeyRequestDto)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		// how come?
		w.WriteHeader(http.StatusInternalServerError)
	}
}
