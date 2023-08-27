package usecases

import (
	"fmt"
	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/dtos"
	"rhSystem_server/app/domain/appointments/services"
	datetimeValue "rhSystem_server/app/domain/appointments/valueobjects/datetime"
	weekdayValue "rhSystem_server/app/domain/appointments/valueobjects/weekday"
	appointmentsRepo "rhSystem_server/app/infrastructure/repositories/appointments"
	"rhSystem_server/app/infrastructure/repositories/interfaces"
	slotsRepo "rhSystem_server/app/infrastructure/repositories/slots"
	"strconv"
	"time"
)

func GetAvaiableSlotsUseCase(avaiableSlotsDTO *dtos.AvaiableSlotsRequestDTO) ([]int, *shared.AppError) {

	// get all valid slots for specific date
	year, _ := strconv.Atoi(avaiableSlotsDTO.Year)
	month, _ := strconv.Atoi(avaiableSlotsDTO.Month)
	day, _ := strconv.Atoi(avaiableSlotsDTO.Day)

	w, err := weekdayValue.New(year, time.Month(month), day)

	if err != nil {
		return nil, err
	}

	var slotsRepository interfaces.SlotsRepositoryInterface
	slotsRepository = slotsRepo.New()

	// w.Value corresponds to the id value but the right way to do it is to fetch the weekday by value and retrieve its id
	validDaySlots, err := services.GetWeekDayValidSlotsService(w.Value, slotsRepository)

	if err != nil {
		return nil, err
	}

	fmt.Println("validDaySlots: ", validDaySlots)
	fmt.Println("len(validDaySlots): ", len(validDaySlots))

	if len(validDaySlots) == 0 {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var appointmentsRepository interfaces.AppointmentsRepositoryInterface
	appointmentsRepository = appointmentsRepo.New()

	datetime, err := datetimeValue.New(year, time.Month(month), day, 0)

	blockedDaySlots, err := services.FilterAppointmentsByDatetimeService(datetime.Value, appointmentsRepository)

	fmt.Println("blockedDaySlots: ", blockedDaySlots)

	if err != nil {
		return nil, err
	}

	// filter validSlots against taken slots ('confirmed', 'pendging' 25min + )
	var filteredSlots []int
	for _, validSlot := range validDaySlots {
		isValid := true
		for _, blockedSlot := range blockedDaySlots {
			if validSlot.Slot == blockedSlot {
				isValid = false
			}
		}

		if isValid {
			filteredSlots = append(filteredSlots, validSlot.Slot)
		}
	}

	fmt.Println("filteredSlots: ", filteredSlots)

	return filteredSlots, nil
}
