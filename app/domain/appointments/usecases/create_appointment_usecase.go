package usecases

import (
	"fmt"
	"net/http"
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/entities"
	"rhSystem_server/app/domain/appointments/services"
	"rhSystem_server/app/infrastructure/database/enums"

	//"rhSystem_server/app/helpers"
	appointmentsRepository "rhSystem_server/app/infrastructure/repositories/appointments"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
	slotsRepository "rhSystem_server/app/infrastructure/repositories/slots"
	"time"

	"github.com/google/uuid"
)

func CreateAppointmentUseCase(newAppointmentDTO *dtos.NewAppointmentRequestDTO) (bool, *shared.AppError) { // status, message, error boolean

	// verify if candidate has pending appointment
	var appointmentsRepo interfaces.AppointmentsRepositoryInterface
	appointmentsRepo = appointmentsRepository.New()
	appointmentFound, err := services.CheckIfCandidateHasAppointmentAlready(newAppointmentDTO.Email, appointmentsRepo)

	if err != nil {
		return false, err
	}

	var channel chan bool
	switch appointmentFound["status"] {
	case enums.Pending:
		// verificar se data de criação do appointment + 25 minutos (confirmation deadline) já passou, se sim update appointment status = 'canceled'
		if v, ok := appointmentFound["created_at"].(time.Time); ok {
			confirmationDeadline := v.Add(25 * time.Minute)

			now := time.Now()
			if now.After(confirmationDeadline) {
				if appointmentId, ok := appointmentFound["id"].(int); ok {
					defer func() {
						channel <- true
					}()

					//services.UpdateAppointmentStatusOnRoutineService(appointmentId, enums.Canceled, appointmentsRepo, channel)

					services.UpdateAppointmentStatusOnRoutineService(
						dtos.UpdateAppointmentStatusDTO{Id: appointmentId, Status: enums.Canceled, Repo: appointmentsRepo, C: channel},
					)
				}
			} else {
				return false, &shared.AppError{Message: "Candidato possui agendamento pendente de confirmação", StatusCode: http.StatusOK}
			}
		} else {
			return false, &shared.AppError{Message: "Candidato já existe", StatusCode: http.StatusInternalServerError}

		}

	case enums.Confirmed:
		// verify if 'confirmed' wich means candidate already did the interview or if some reason they didn't show must have a
		// update cancel operation then
		return false, &shared.AppError{Message: "Candidato já realizou entrevista", StatusCode: http.StatusBadRequest}
	case enums.Canceled:
		// if canceled proceed to appointment scheduling
		fmt.Println("Canceled")
	default:
		fmt.Println("Default")
	}

	var slotRepo interfaces.SlotsRepositoryInterface
	slotRepo = slotsRepository.New()
	slotAvaiable, err := services.CheckIfSlotAvaiable(newAppointmentDTO.Slot, slotRepo)

	if err != nil {
		return false, err
	}

	if !slotAvaiable {
		return false, &shared.AppError{Message: "Slot inexistente", StatusCode: http.StatusBadRequest}
	}

	// iniciar transaction (candidate & appointment)
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

	_, err = services.CreateAppointmentOnTransactionService(&newCandidate, &newAppointment)

	if err != nil {
		return false, err
	}

	fmt.Println("SUCCESS: Call defer channel func!")

	// needs to have a goroutine to receive the value otherwise execution will stay blocked (use defer maybe)
	//channel <- true

	go func() {
		fmt.Println("Send email!")
		//services.SendConfirmationEmail(newCandidate.Email)
	}()

	return true, nil
}
