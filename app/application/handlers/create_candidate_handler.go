package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rhSystem_server/app/domain/candidates/dtos"
	"rhSystem_server/app/domain/candidates/services"
	"rhSystem_server/app/infrastructure/repositories/interfaces"

	repositories "rhSystem_server/app/infrastructure/repositories/candidates"
)

func CreateCandidateHandler(w http.ResponseWriter, r *http.Request) {

	var newCandidateDto dtos.CandidateRequestDTO

	decodeErr := json.NewDecoder(r.Body).Decode(&newCandidateDto)

	if decodeErr != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	fmt.Println("newCandidateDto: ", newCandidateDto.Name)

	var repo interfaces.CandidateRepositoryInterface
	repo = repositories.New()
	createdCandidate, err := services.CreateCandidateService(&newCandidateDto, repo)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	encodeErr := json.NewEncoder(w).Encode(&createdCandidate)

	if encodeErr != nil {
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		return
	}
}
