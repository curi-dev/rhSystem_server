package repositories

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/lib/pq"

	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/candidates/entities"
	valueobjects "rhSystem_server/app/domain/candidates/valueobjects/accessKey"
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

func (repository *PostgresCandidatesRepository) FindByEmail(email string) (*entities.Candidate, *shared.AppError) { // if slot is valid or not

	rows, err := repository.db.Query(
		`SELECT id, name, phone FROM candidates WHERE email = $1`,
		email,
	)

	fmt.Println("rows: ", rows)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema no servidor", StatusCode: http.StatusInternalServerError}
	}

	defer rows.Close()

	if rows.Next() {
		var id uuid.UUID
		var name string
		var phone string

		if err := rows.Scan(&id, &name, &phone); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema no servidor", StatusCode: http.StatusInternalServerError}
		}

		if err := rows.Err(); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema no servidor", StatusCode: http.StatusInternalServerError}
		}

		candidate := entities.Candidate{Id: id, Name: name, Email: email, Phone: phone}

		return &candidate, nil
	}

	return nil, nil
}

func (repository *PostgresCandidatesRepository) AccessKey(k *valueobjects.AccessKey) (*valueobjects.AccessKey, *shared.AppError) {

	result, err := repository.db.Exec(
		`INSERT INTO "access_keys" (id, value, candidate) VALUES ($1, $2, $3)`,
		k.Id,
		k.Value,
		k.Candidate,
	)

	if err != nil {
		fmt.Println("err: ", err)

		if pgError, ok := err.(*pq.Error); ok {
			if pgError.Code == "23505" {
				return nil, &shared.AppError{Err: pgError, Message: "Ocorreu um problema no servicdor", StatusCode: http.StatusBadRequest}
			}
		}

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema no servidor", StatusCode: http.StatusInternalServerError}
	}

	rowsAffected, err := result.RowsAffected()

	fmt.Println("rowsAffected: ", rowsAffected)

	if err != nil {
		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	// returns the same data
	return k, nil
}
