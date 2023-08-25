package repositories

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lib/pq"

	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/entities"
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

func (repository *PostgresCandidatesRepository) Create(c *entities.Candidate) (*entities.Candidate, *shared.AppError) {

	_, err := repository.db.Exec(
		`INSERT INTO candidates (id, name, email, phone) VALUES ($1, $2, $3, $4)`,
		c.Id,
		c.Name,
		c.Email,
		c.Phone,
	)

	if err != nil {

		fmt.Println("err: ", err)

		if pgError, ok := err.(*pq.Error); ok {
			if pgError.Code == "23505" {
				return nil, &shared.AppError{Err: pgError, Message: "Candidato já existe", StatusCode: http.StatusBadRequest}
			}
		}

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	// returns the same data
	return c, nil
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
