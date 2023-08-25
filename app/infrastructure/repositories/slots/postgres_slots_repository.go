package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/database"
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
