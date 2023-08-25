package handlers

import (
	"encoding/json"
	"net/http"
	"rhSystem_server/app/domain/candidates/entities"
	"rhSystem_server/app/domain/candidates/services"
	repositories "rhSystem_server/app/infrastructure/repositories/candidates"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
)

func CreateCandidateHandler(w http.ResponseWriter, r *http.Request) {

	var newCandidate entities.Candidate

	decodeErr := json.NewDecoder(r.Body).Decode(&newCandidate)

	if decodeErr != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	var repo interfaces.CandidateRepositoryInterface
	repo = repositories.New()
	createdCandidate, err := services.CreateCandidateService(&newCandidate, repo)

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
