package usecases

import (
	"fmt"
	"net/http"
	"os"
	shared "rhSystem_server/app/application/error"
	applicationServices "rhSystem_server/app/application/services"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/entities"
	"rhSystem_server/app/domain/appointments/services"
	"strconv"

	// validSlot "rhSystem_server/app/domain/appointments/valueobjects/validSlot"
	datetime "rhSystem_server/app/domain/appointments/valueobjects/datetime"
	hour "rhSystem_server/app/domain/appointments/valueobjects/hour"
	"rhSystem_server/app/infrastructure/database/enums"

	//"rhSystem_server/app/helpers"
	appointmentsRepository "rhSystem_server/app/infrastructure/repositories/appointments"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
	slotsRepository "rhSystem_server/app/infrastructure/repositories/slots"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func CreateAppointmentUseCase(newAppointmentDTO *dtos.AppointmentRequestDTO) (bool, *shared.AppError) { // status, message, error boolean

	var channel chan bool

	var appointmentsRepo interfaces.AppointmentsRepositoryInterface
	appointmentsRepo = appointmentsRepository.New()
	appointmentFound, err := services.CheckIfCandidateHasAppointmentAlready(newAppointmentDTO.CandidateId, appointmentsRepo)

	if err != nil {
		return false, err
	}
	switch appointmentFound["status"] {
	case enums.Pending:
		// verificar se data de criação do appointment + 25 minutos (confirmation deadline) já passou, se sim update appointment status = 'canceled'
		if v, ok := appointmentFound["created_at"].(time.Time); ok {

			elapsedTimeFromSchedule := time.Since(v)
			if elapsedTimeFromSchedule.Minutes()-3*60 > 25 { // timezone compensation
				if appointmentId, ok := appointmentFound["id"].(int); ok {

					// if transaction does not commit execute/update appointment to 'canceled' on defer mode anyway
					defer func() {
						channel <- true
					}()

					services.UpdateAppointmentStatusOnRoutineService(
						dtos.UpdateAppointmentStatusDTO{Id: appointmentId, Status: enums.Canceled, Repo: appointmentsRepo, C: channel},
					)
				}
			} else {
				return false, &shared.AppError{Message: "Candidato possui agendamento pendente de confirmação", StatusCode: http.StatusPreconditionFailed}
			}
		} else {
			return false, &shared.AppError{Message: "Ocorreu um erro no servidor", StatusCode: http.StatusInternalServerError}
		}

	case enums.Confirmed:
		// verify if 'confirmed' wich means candidate already did the interview or if some reason they didn't show must have a
		// update cancel operation then
		return false, &shared.AppError{Message: "Candidato já agendou entrevista", StatusCode: http.StatusBadRequest}
	case enums.Canceled:
		// if canceled proceed to appointment scheduling
		fmt.Println("Canceled")
	default:
		fmt.Println("Default")
	}

	var slotRepo interfaces.SlotsRepositoryInterface
	slotRepo = slotsRepository.New()
	slotAvaiable, err := services.CheckIfSlotExists(newAppointmentDTO.Slot, &newAppointmentDTO.SplittedDate, slotRepo)

	if err != nil {
		return false, err
	}

	if !slotAvaiable {
		return false, &shared.AppError{Message: "Slot inexistente", StatusCode: http.StatusBadRequest}
	}

	hour, constructorErr := hour.New(newAppointmentDTO.Slot)
	if constructorErr != nil {
		return false, constructorErr
	}

	datetime, constructorErr := datetime.New(
		newAppointmentDTO.SplittedDate.Year,
		newAppointmentDTO.SplittedDate.Month,
		newAppointmentDTO.SplittedDate.Day,
		hour.Value,
	)

	if constructorErr != nil {
		return false, constructorErr
	}

	var newAppointment entities.Appointment
	newAppointment.Candidate = newAppointmentDTO.CandidateId
	newAppointment.Status = enums.Pending
	newAppointment.Id = uuid.New()
	newAppointment.Datetime = datetime.Value
	newAppointment.Slot = newAppointmentDTO.Slot

	success, err := services.CreateAppointmentService(&newAppointment, newAppointmentDTO.CandidateId.String(), appointmentsRepo)

	if err != nil {
		return false, err
	}

	// needs to have a goroutine to receive the value otherwise execution will stay blocked (use defer maybe)
	// channel <- true
	fmt.Println("Success: Call defer channel func: ", success)

	if success {
		go func() {

			godotenv.Load("env")

			fmt.Println("Send email!")

			slotString := strconv.Itoa(newAppointment.Slot)
			dayString := strconv.Itoa(newAppointmentDTO.SplittedDate.Day)
			yearString := strconv.Itoa(newAppointmentDTO.SplittedDate.Year)
			monthString := strconv.Itoa(int(newAppointmentDTO.SplittedDate.Month))

			confirmationLink := os.Getenv("CONFIRMATION_LINK")
			subject := "Subject: Link de confirmação\n\n"
			body := fmt.Sprintf("%s?apnmnt=%s&slot=%s&day=%s&month=%s&year=%s",
				confirmationLink,
				newAppointment.Id.String(),
				slotString,
				dayString,
				monthString,
				yearString,
			)

			err := applicationServices.SendEmail(newAppointmentDTO.CandidateEmail, subject, body) // id is coming from the prebuilt struct (before databse insertion)

			if err != nil {
				fmt.Println("error: ", err.Error())
			}
		}()
	}

	return success, nil
}
