package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/database"
)

type PostgresAppointmentsRepository struct {
	db *sql.DB
}

func New() *PostgresAppointmentsRepository {
	return &PostgresAppointmentsRepository{
		db: database.GetDB(),
	}
}

func (repository *PostgresAppointmentsRepository) FindByCandidateEmail(email string) (map[string]interface{}, *shared.AppError) { // if slot is valid or not

	rows, err := repository.db.Query(
		`SELECT id, status, created_at FROM appointments WHERE email = $1`,
		email,
	)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	defer rows.Close()

	if rows.Next() {
		var id string
		var status int
		var createdAt time.Time
		if err := rows.Scan(
			&id,
			&status,
			&createdAt,
		); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		if err := rows.Err(); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		response := make(map[string]interface{})
		response["id"] = id
		response["status"] = status
		response["created_at"] = createdAt

		fmt.Println("response: ", response)

		return response, nil
	}

	// no appointment from that candidate found and no error
	return nil, nil
}

func (repository *PostgresAppointmentsRepository) UpdateStatus(id int, status int) (bool, *shared.AppError) { // if slot is valid or not

	result, err := repository.db.Exec(
		`UPDATE appointments SET status = $1 WHERE id = $2`,
		status,
		id,
	)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	if affected > 0 {
		return true, nil
	}

	return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
}
