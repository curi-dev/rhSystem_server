package repositories

import (
	"database/sql"
	"fmt"
	"net/http"

	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/infrastructure/database"
)

type PostgresCandidatesRepository struct {
	db *sql.DB
}

func New() *PostgresCandidatesRepository {
	return &PostgresCandidatesRepository{
		db: database.GetDB(),
	}
}

func (repository *PostgresCandidatesRepository) FindByEmail(email string) (bool, *shared.AppError) { // if slot is valid or not

	rows, err := repository.db.Query(
		`SELECT id FROM candidates WHERE email = $1`,
		email,
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
