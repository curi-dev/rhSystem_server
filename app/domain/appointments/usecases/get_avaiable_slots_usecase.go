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
)

func GetAvaiableSlotsUseCase(avaiableSlotsDTO *dtos.AvaiableSlotsRequestDTO) ([]int, *shared.AppError) {

	// get all valid slots for specific date
	w, err := weekdayValue.New(avaiableSlotsDTO.SplittedDate.Year, avaiableSlotsDTO.SplittedDate.Month, avaiableSlotsDTO.SplittedDate.Day)

	if err != nil {
		return nil, err
	}

	var slotsRepository interfaces.SlotsRepositoryInterface
	slotsRepository = slotsRepo.New()

	// it represents the avaiable slot values for a specific day
	var avaiableSlots []int
	validDaySlots, err := services.GetWeekDayValidSlotsService(w.Value, slotsRepository)

	if len(validDaySlots) == 0 {
		return avaiableSlots, nil
	}

	if err != nil {
		return nil, err
	}

	var appointmentsRepository interfaces.AppointmentsRepositoryInterface
	appointmentsRepository = appointmentsRepo.New()

	datetime, err := datetimeValue.New(avaiableSlotsDTO.SplittedDate.Year, avaiableSlotsDTO.SplittedDate.Month, avaiableSlotsDTO.SplittedDate.Day, 0)

	blockedDaySlots, err := services.FilterAppointmentsByDatetimeService(datetime.Value, appointmentsRepository)

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
