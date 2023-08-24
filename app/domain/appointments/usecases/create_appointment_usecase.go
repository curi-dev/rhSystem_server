package usecases

import (
	"net/http"
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/entities"
	"rhSystem_server/app/domain/appointments/services"
	"rhSystem_server/app/infrastructure/database/enums"
	candidatesRepository "rhSystem_server/app/infrastructure/repositories/candidates"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
	slotsRepository "rhSystem_server/app/infrastructure/repositories/slots"

	"github.com/google/uuid"
)

func CreateAppointmentUseCase(newAppointmentDTO *dtos.NewAppointmentRequestDTO) (bool, *shared.AppError) { // status, message, error boolean

	var repo interfaces.CandidateRepositoryInterface
	repo = candidatesRepository.New()
	candidateExists, err := services.CheckIfCandidateExistsService(newAppointmentDTO.Email, repo)

	if err != nil {
		return false, err
	}

	if candidateExists {
		return false, &shared.AppError{Message: "Usuário já existe", StatusCode: http.StatusBadRequest}
	}

	// VERIFICA SE O USUÁRIO POSSUI APPOINTMENT PENDING

	// newAppointmentDTO contém o id do slot que o usuário selecionou
	var slotRepo interfaces.SlotsRepositoryInterface
	slotRepo = slotsRepository.New()
	slotAvaiable, err := services.CheckIfSlotAvaiable(newAppointmentDTO.Slot, slotRepo)

	if err != nil {
		return false, err
	}

	if !slotAvaiable {
		return false, &shared.AppError{Message: "Slot inexistente", StatusCode: http.StatusInternalServerError}
	}

	// CRIAR TRANSAÇÃO PARA CANDIDATOS E APPOINTMENTS
	newCandidate := entities.Candidate{
		Id:    uuid.New(),
		Name:  newAppointmentDTO.Name,
		Email: newAppointmentDTO.Email,
		Phone: newAppointmentDTO.Phone,
	}

	var newAppointment entities.Appointment
	newAppointment.Candidate = newCandidate.Id
	newAppointment.Status = enums.Pending
	newAppointment.Id = uuid.New()
	newAppointment.Datetime = newAppointmentDTO.Datetime
	newAppointment.Slot = newAppointmentDTO.Slot

	success := services.CreateAppointmentOnTransactionService(&newCandidate, &newAppointment)

	if !success {
		return false, err
	}

	// -> do it asyncronously
	//services.SendConfirmationEmail(), nil

	return true, nil
}
