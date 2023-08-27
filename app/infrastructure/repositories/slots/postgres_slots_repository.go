package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	shared "rhSystem_server/app/application/error"
	slotAggregate "rhSystem_server/app/domain/appointments/aggregate/slot"
	validslotAggregate "rhSystem_server/app/domain/appointments/aggregate/validSlot"
	"rhSystem_server/app/infrastructure/database"

	"github.com/google/uuid"
	//"github.com/google/uuid"
	//"github.com/lib/pq"
	//"rhSystem_server/app/domain/appointments/entities"
	//"rhSystem_server/app/domain/appointments/valueobjects"
	//"rhSystem_server/app/domain/appoitments/valueobjects"
)

type PostgresSlotsRepository struct {
	db *sql.DB
}

func New() *PostgresSlotsRepository {
	return &PostgresSlotsRepository{
		db: database.GetDB(),
	}
}

func (repository *PostgresSlotsRepository) Index() ([]slotAggregate.Slot, *shared.AppError) {

	rows, err := repository.db.Query("SELECT id, label, value FROM slots")

	defer rows.Close()

	fmt.Println("rows: ", rows)

	if err != nil {
		if err != nil {
			fmt.Println("Error during query, ", err)

			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}
	}

	var slots []slotAggregate.Slot

	for rows.Next() {
		fmt.Println("Next()")

		var id string
		var label string
		var value string
		if err := rows.Scan(
			&id,
			&label,
			&value,
		); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		slots = append(slots, slotAggregate.Slot{Id: id, Label: label, Value: value})
	}

	if err := rows.Err(); err != nil {
		return slots, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	return slots, nil
}

func (repository *PostgresSlotsRepository) FindByWeekday(weekdayId int) ([]validslotAggregate.ValidSlot, *shared.AppError) {

	fmt.Println("weekdayId: ", weekdayId)

	query := `SELECT id, weekday, slot FROM "valid_slots" WHERE weekday = $1`
	rows, err := repository.db.Query(
		query,
		weekdayId,
	)

	fmt.Println("Query: ", query)

	defer rows.Close()

	if err != nil {
		fmt.Println("Error during query, ", err)

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	var validSlots []validslotAggregate.ValidSlot

	for rows.Next() {

		var (
			id      uuid.UUID
			weekday int
			slot    int
		)

		if err := rows.Scan(
			&id,
			&weekday,
			&slot,
		); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		validSlots = append(validSlots, validslotAggregate.ValidSlot{Id: id, Weekday: weekday, Slot: slot})
	}

	if err := rows.Err(); err != nil {
		return validSlots, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	return validSlots, nil
}

func (repository *PostgresSlotsRepository) Find(w *time.Weekday, slot int) (bool, *shared.AppError) { // if slot is valid or not

	rows, err := repository.db.Query(
		`SELECT id, weekday, slot FROM "valid_slots" WHERE weekday = $1 AND slot = $2`,
		w,
		slot,
	)

	fmt.Println("rows: ", rows)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.Err(); err != nil {
			return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		return true, nil
	}

	return false, nil
}
